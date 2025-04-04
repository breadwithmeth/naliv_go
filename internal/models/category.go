package models

type Category struct {
	ID                 int     `json:"category_id"`
	Name               string  `json:"name"`
	ParentCategory     *int    `json:"parent_category"`
	ParentCategoryName *string `json:"parent_category_name"`
}
