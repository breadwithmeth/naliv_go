package models

type User struct {
	ID   int     `json:"user_id"`
	Name *string `json:"name,omitempty"`
}
