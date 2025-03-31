package models

import "time"

type Item struct {
	ID         int             `json:"item_id"`
	Name       string          `json:"name"`
	Code       string          `json:"code"`
	Image      string          `json:"img"`
	Category   Category        `json:"category"`
	Price      float64         `json:"price"`
	InStock    float64         `json:"in_stock"`
	Unit       string          `json:"unit"`
	Quantity   float64         `json:"quantity"`
	Options    []ItemOption    `json:"options"`
	Promotions []ItemPromotion `json:"promotions"`
}

type ItemOption struct {
	Selection string       `json:"selection"`
	Required  bool         `json:"required"`
	Name      string       `json:"name"`
	Options   []OptionItem `json:"options"`
}

type OptionItem struct {
	Price            float64 `json:"price"`
	RelationID       int     `json:"relation_id"`
	Name             string  `json:"name"`
	ParentItemAmount float64 `json:"parent_item_amount"`
}

type ItemPromotion struct {
	EndDate    time.Time `json:"end_date"`
	StartDate  time.Time `json:"start_date"`
	BaseAmount float64   `json:"base_amount"`
	AddAmount  float64   `json:"add_amount"`
	Name       string    `json:"name"`
}
