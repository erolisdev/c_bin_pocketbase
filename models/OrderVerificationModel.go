package models

// Product represents a product in the order0
// type Product struct {
// 	ProductID    string        `json:"product_id"`
// 	Name         string        `json:"name"`
// 	Category     string        `json:"category"`
// 	Price        float64       `json:"price"` // Price PER UNIT of the product
// 	Quantity     int           `json:"quantity"`
// 	OptionValues []OptionValue `json:"options"`
// }

// // OrderInfo represents order metadata
// type OrderInfo struct {
// 	OrderDate  string  `json:"order_date"`
// 	TotalPrice float64 `json:"total_price"`
// 	Currency   string  `json:"currency"`
// 	Status     string  `json:"status"`
// }

// // Order represents the complete order
// type Order struct {
// 	OrderInfo     OrderInfo `json:"order_info"`
// 	OrderProducts []Product `json:"order_products"`
// }

// // OptionValue represents option value from database or order
// type OptionValue struct {
// 	ID            string  `json:"id,omitempty"`         // DB ID of the option_values record
// 	OptionID      string  `json:"option_id"`            // ID of the parent option (e.g., "size", "topping")
// 	Name          string  `json:"name,omitempty"`       // Name of the option value (e.g., "Large", "Pepperoni")
// 	Quantity      int     `json:"quantity"`             // How many of this specific option value are selected
// 	Price         float64 `json:"price"`                // Price PER UNIT of this option value AS SENT BY CLIENT
// 	FreeCount     int     `json:"free_count,omitempty"` // From DB options table (informational, not used for client data)
// 	MaxCount      int     `json:"max_count,omitempty"`  // From DB options table (informational, not used for client data)
// 	OptionValueID string  `json:"option_value_id"`      // Specific ID for this value (e.g., "pepperoni_id", "large_size_id")
// 	OptionType    string  `json:"type,omitempty"`       // From DB options table (informational, not used for client data)
// }

// // DBDataOptionValue is used to store data fetched from the DB for option values
// type DBDataOptionValue struct {
// 	ID            string  // Record ID from option_values table
// 	OptionID      string  // FK to options table
// 	Name          string  // Name from options table (for the group)
// 	PriceTTC      float64 // Price from option_values table
// 	FreeCount     int     // From options table
// 	MaxCount      int     // From options table
// 	OptionType    string  // From options table
// 	OptionValueID string  // The unique ID for this value (e.g. "red_color_id")
// }
