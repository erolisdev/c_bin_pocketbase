package models

type StoreProduct struct {
	ProductID       int           `json:"product_id,omitempty"`
	CategoryID      int           `json:"category_id,omitempty"`
	Image           string        `json:"image,omitempty"`
	Shipping        int           `json:"shipping,omitempty"`
	Price           string        `json:"price,omitempty"`
	PrinterID       int           `json:"printer_id,omitempty"`
	LblPrinterID    int           `json:"lbl_printer_id"`
	SortOrder       int           `json:"sort_order,omitempty"`
	ShowInSuggested bool          `json:"show_in_suggested,omitempty"`
	Rate            string        `json:"rate,omitempty"`
	Descriptions    any           `json:"description,omitempty"`
	Options         []StoreOption `json:"options,omitempty"`
	TaxRate         string        `json:"tax_rate,omitempty"`
}
