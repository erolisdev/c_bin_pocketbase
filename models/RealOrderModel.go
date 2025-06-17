package models

const Radio = "radio"
const General = "general"
const Special = "special"
const OneIsFree = "one_is_free"

type OrderModel struct {
	OrderData  *OrderData `json:"items,omitempty"`
	SuccessURL *string    `json:"success_url,omitempty"`
	CancelURL  *string    `json:"cancel_url,omitempty"`
}

type OrderData struct {
	Reference            *string        `json:"reference,omitempty"`
	StoreID              *int64         `json:"store_id,omitempty"`
	StoreURL             *string        `json:"store_url,omitempty"`
	StoreName            *string        `json:"store_name,omitempty"`
	DeliveryTime         *string        `json:"delivery_time,omitempty"`
	OrderNumber          *int           `json:"order_number,omitempty"`
	OrderTime            *string        `json:"order_time"`
	OrderStatus          *int64         `json:"order_status,omitempty"`
	TableNo              *string        `json:"table_no"`
	CustomerID           *int64         `json:"customer_id,omitempty"`
	CustomerFirstname    *string        `json:"customer_firstname,omitempty"`
	CustomerLastname     *string        `json:"customer_lastname,omitempty"`
	CustomerEmail        *string        `json:"customer_email,omitempty"`
	CustomerTelephone    *string        `json:"customer_telephone,omitempty"`
	CustomerDeliveryTime *string        `json:"customer_delivery_time,omitempty"`
	PaymentMethod        *string        `json:"payment_method,omitempty"`
	PaymentCode          *string        `json:"payment_code,omitempty"`
	ShippingFirstname    *string        `json:"shipping_firstname,omitempty"`
	ShippingLastname     *string        `json:"shipping_lastname,omitempty"`
	ShippingCompany      *string        `json:"shipping_company"`
	ShippingAddress1     *string        `json:"shipping_address_1"`
	ShippingAddress2     *string        `json:"shipping_address_2"`
	ShippingCity         *string        `json:"shipping_city"`
	ShippingPostcode     *string        `json:"shipping_postcode"`
	ShippingCityID       *string        `json:"shipping_city_id"`
	ShippingCountryID    *string        `json:"shipping_country_id"`
	ShippingZoneID       *string        `json:"shipping_zone_id"`
	ShippingMethod       *string        `json:"shipping_method,omitempty"`
	ShippingCode         *string        `json:"shipping_code,omitempty"`
	ShippingPrice        *string        `json:"shipping_price,omitempty"`
	Comment              *string        `json:"comment"`
	Total                *string        `json:"total,omitempty"`
	TotalTax             *string        `json:"total_tax,omitempty"`
	AffiliateID          *int64         `json:"affiliate_id,omitempty"`
	Commission           *string        `json:"commission,omitempty"`
	MarketingID          *int64         `json:"marketing_id,omitempty"`
	Tracking             *string        `json:"tracking,omitempty"`
	LanguageID           *int64         `json:"language_id,omitempty"`
	CurrencyID           *int64         `json:"currency_id,omitempty"`
	CurrencyCode         *string        `json:"currency_code,omitempty"`
	IP                   *string        `json:"ip,omitempty"`
	ForwaredIP           *string        `json:"forwared_ip,omitempty"`
	UserAgent            *string        `json:"user_agent,omitempty"`
	AcceptLanguage       *string        `json:"accept_language,omitempty"`
	PrinterID            *int64         `json:"printer_id,omitempty"`
	PrinterPort          *int64         `json:"printer_port,omitempty"`
	PrinterIP            *string        `json:"printer_ip,omitempty"`
	TemplateID           *int64         `json:"template_id,omitempty"`
	TokenCAPTCHA         *string        `json:"token_captcha"`
	OrderProducts        []OrderProduct `json:"products,omitempty"`
}

type OrderProduct struct {
	LblPrinterID *int64             `json:"lbl_printer_id"`
	CategoryID   *int64             `json:"category_id,omitempty"`
	CSortOrder   *int64             `json:"c_sort_order,omitempty"`
	PrinterPort  *int64             `json:"printer_port,omitempty"`
	PrinterID    *int64             `json:"printer_id,omitempty"`
	ProductID    *int64             `json:"product_id,omitempty"`
	Name         *string            `json:"name,omitempty"`
	Quantity     *int64             `json:"quantity,omitempty"`
	Price        *string            `json:"price,omitempty"`
	Total        *string            `json:"total,omitempty"`
	Tax          *string            `json:"tax,omitempty"`
	TaxRate      *string            `json:"tax_rate,omitempty"`
	OptionValues []OrderOptionValue `json:"option,omitempty"`
}

type OrderOptionValue struct {
	ProductOptionID      *int64   `json:"product_option_id,omitempty"`
	ProductOptionValueID *int64   `json:"product_option_value_id,omitempty"`
	OptionID             *int64   `json:"option_id,omitempty"`
	OptionValueID        *int64   `json:"option_value_id,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	Value                *string  `json:"value,omitempty"`
	Type                 *string  `json:"type,omitempty"`
	Quantity             *float64 `json:"quantity,omitempty"` // How many of this specific option value are selected
	Price                *string  `json:"price,omitempty"`
	PriceHt              *string  `json:"price_ht,omitempty"`
	TotalTTC             *string  `json:"total_ttc,omitempty"`
	TotalHT              *string  `json:"total_ht,omitempty"`
	Tax                  *string  `json:"tax,omitempty"`
	TaxRate              *string  `json:"tax_rate,omitempty"`
	Title                *int64   `json:"title,omitempty"`
	PoSortOrder          *int64   `json:"po_sort_order,omitempty"`
}

// DBDataOptionValue is used to store data fetched from the DB for option values
type DBDataOptionValue struct {
	ID            string  // Record ID from option_values table
	OptionID      int64   // FK to options table
	Name          string  // Name from options table (for the group)
	PriceTTC      float64 // Price from option_values table
	FreeCount     int     // From options table
	MaxCount      int     // From options table
	OptionType    string  // From options table
	PriceStatus   string  // From options table
	OptionValueID int64   // The unique ID for this value (e.g. "red_color_id")
}
