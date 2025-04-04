package models

type Address struct {
	ID         int     `json:"address_id"`
	Address    *string `json:"address,omitempty"`
	Apartment  *string `json:"apartment,omitempty"`
	Entrance   *string `json:"entrance,omitempty"`
	Floor      *string `json:"floor,omitempty"`
	CityName   *string `json:"city_name,omitempty"`
	IsSelected *int    `json:"is_selected,omitempty"`
}
