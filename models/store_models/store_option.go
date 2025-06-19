package models

type StoreOption struct {
	OptionID        int                 `db:"option_id" json:"option_id,omitempty"`
	PoSortOrder     int                 `json:"po_sort_order,omitempty"`
	Required        int                 `db:"required" json:"required,omitempty"`
	FreeOptionCount int                 `db:"free_option_count" json:"free_option_count,omitempty"`
	MaxOptionCount  int                 `db:"max_option_count" json:"max_option_count,omitempty"`
	Type            string              `db:"type" json:"type,omitempty"`
	OptionGroup     int                 `db:"option_group" json:"option_group,omitempty"`
	Title           int                 `db:"title" json:"title,omitempty"`
	Descriptions    any                 `db:"descriptions" json:"language,omitempty"`
	Values          *[]StoreOptionValue `json:"values,omitempty"`
}

type StoreOptionValue struct {
	OptionID        int    `db:"option_id" json:"option_id,omitempty"`
	OptionValueID   int    `db:"option_value_id" json:"option_value_id,omitempty"`
	Image           string `db:"image_url" json:"image,omitempty"`
	SortOrder       int    `db:"sort_order" json:"sort_order,omitempty"`
	RelatedOptionID any    `db:"related_option_id"  json:"related_option_id,omitempty"`
	PriceStatus     string `db:"price_status" json:"price_status,omitempty"`
	Reset           int    `db:"reset" json:"reset,omitempty"`
	Grup            string `db:"grup" json:"grup,omitempty"`
	ProductID       *int   `json:"product_id,omitempty"`
	PriceTTC        string `db:"price_ttc" json:"ov_price,omitempty"`
	PriceHT         string `db:"price_ht" json:"ov_price_ht,omitempty"`
	Descriptions    any    `db:"descriptions" json:"language,omitempty"`
}
