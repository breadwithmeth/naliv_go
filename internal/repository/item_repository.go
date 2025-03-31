package repository

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/breadwithmeth/naliv_go/internal/models"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &ItemRepository{db: db}
}

type GetItemsParams struct {
	BusinessID int
	CategoryID *int
}

func (r *ItemRepository) GetItems(business_id int) ([]*models.Item, error) {
	query := `
	SELECT JSON_OBJECT(
		'items', JSON_ARRAYAGG(
			JSON_OBJECT(
				'name', ii.name,
				'code', ii.code,
				'img', ii.img,
				'item_id', ii.item_id,
				'category', (SELECT JSON_OBJECT(
					'category_id', c.category_id,
					'name', c.name,
					'parent_category', c.parent_category,
					'parent_category_name', c2.name
				) FROM categories c LEFT JOIN categories c2 ON c2.category_id = c.parent_category WHERE c.category_id = ii.category_id),
				'options', (
					SELECT JSON_ARRAYAGG(
						JSON_OBJECT(
							'selection', options.selection,
							'required', IF(options.required = 1, cast(TRUE as json), cast(FALSE as json)),
							'name', options.name,
							'options', (
								SELECT JSON_ARRAYAGG(
									JSON_OBJECT(
										'price', option_items.price,
										'relation_id', option_items.relation_id,
										'name', i3.name,
										'parent_item_amount', option_items.parent_item_amount
									)
								)
								FROM option_items
								LEFT JOIN items i3 ON i3.item_id = option_items.item_id
								WHERE option_items.option_id = options.option_id
							)
						)
					)
					FROM options 
					WHERE options.item_id = ii.item_id
				),
				'price', FLOOR(ii.price),
				'in_stock', ii.in_stock,
				'unit', ii.unit,
				'quantity', ii.quantity,
				'promotions', (
					SELECT JSON_ARRAYAGG(
						JSON_OBJECT(
							
							'base_amount', mpd.base_amount,
							'add_amount', mpd.add_amount,
							'name', mpd.name
						)
					)
					FROM marketing_promotion_details mpd
					LEFT JOIN marketing_promotions mp ON mp.marketing_promotion_id = mpd.marketing_promotion_id
					WHERE mpd.item_id = ii.item_id 
					AND CURRENT_TIMESTAMP < mp.end_promotion_date 
					AND CURRENT_TIMESTAMP > mp.start_promotion_date
				)
			)
		)
	) as items
	FROM (
		SELECT i.name, i.code, i.img, i.item_id, i.price, i.amount as in_stock, 
			   i.unit, i.quantity, i.business_id, i.category_id
		FROM items i
		WHERE i.business_id = ? AND i.visible = 1 AND i.price > 1 AND i.amount >= 0`

	// if params.CategoryID != nil {
	// 	query += ` AND i.category_id IN (SELECT category_id FROM categories WHERE parent_category = ?)`
	// }

	query += ` ORDER BY i.price DESC) ii`

	rows, err := r.db.Query(query, business_id)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	if rows.Next() {
		var jsonData string
		if err := rows.Scan(&jsonData); err != nil {
			log.Println("Error scanning JSON data:", err)
			return nil, err
		}

		if jsonData == "" {
			log.Println("Empty JSON data received")
			return []*models.Item{}, nil
		}

		var result struct {
			Items []models.Item `json:"items"`
		}

		if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
			log.Println("Error unmarshalling JSON:", err, "JSON data:", jsonData)
			return nil, err
		}

		for _, item := range result.Items {
			itemCopy := item
			items = append(items, &itemCopy)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return items, nil
}
