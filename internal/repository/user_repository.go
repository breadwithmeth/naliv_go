package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/breadwithmeth/naliv_go/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT user_id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepository) ValidateToken(token string) (int, error) {
	var userID int
	query := "SELECT user_id FROM users_tokens WHERE token = ?"
	err := r.db.QueryRow(query, token).Scan(&userID)
	log.Println("Validating token:", token)
	log.Println("User ID from token:", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("invalid token")
		}
		return 0, err
	}
	return userID, nil
}

func (r *UserRepository) GetUserAddresses(userID int) ([]*models.Address, error) {
	query := `
        SELECT 
            ua.address_id,
            ua.address,
            ua.apartment,
            ua.entrance,
            ua.floor,
            c.name as city_name,
            CASE 
                WHEN sa.address_id IS NOT NULL THEN 1 
                ELSE 0 
            END as is_selected
        FROM user_addreses ua
        LEFT JOIN cities c ON ua.city_id = c.city_id
        LEFT JOIN (
            SELECT address_id 
            FROM selected_address 
            WHERE user_id = ?
            ORDER BY relation_id DESC 
            LIMIT 1
        ) sa ON ua.address_id = sa.address_id
        WHERE ua.user_id = ? AND ua.isDeleted = 0`

	rows, err := r.db.Query(query, userID, userID)
	if err != nil {
		log.Printf("Error querying addresses for user %d: %v", userID, err)
		return nil, err
	}
	defer rows.Close()

	var addresses []*models.Address
	for rows.Next() {
		addr := &models.Address{}
		err := rows.Scan(
			&addr.ID,
			&addr.Address,
			&addr.Apartment,
			&addr.Entrance,
			&addr.Floor,
			&addr.CityName,
			&addr.IsSelected,
		)
		if err != nil {
			log.Printf("Error scanning address row: %v", err)
			return nil, err
		}
		addresses = append(addresses, addr)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after scanning rows: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved %d addresses for user %d", len(addresses), userID)
	return addresses, nil
}
