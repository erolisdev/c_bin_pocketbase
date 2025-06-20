package models

type StoreCategory struct {
	CategoryID      int    `json:"category_id,omitempty"`
	Column          int    `json:"column,omitempty"`
	SortOrder       int    `json:"sort_order,omitempty"`
	Status          int    `json:"status,omitempty"`
	Image           string `json:"image_url,omitempty"`
	ShowInSuggested int    `json:"show_in_suggested,omitempty"`
	Name            string `json:"name,omitempty"`
	Descriptions    any    `json:"description,omitempty"`
}
