package constants

const (
	TableCustomers string = "customers"
	TableAuditLogs string = "audit"

	//store data
	TableStoreCategories   string = "store_categories"
	TableStoreProducts     string = "store_products"
	TableStoreOptions      string = "store_options"
	TableStoreOptionValues string = "store_option_values"
	TableStoreTables       string = "store_tables"
	TableStorePrinters     string = "store_printers"
	TableStoreSettings     string = "store_settings"

	//live orders
	TableLiveOrders              string = "live_orders"
	TableLiveOrderProducts       string = "live_order_products"
	TableLiveOrderProductOptions string = "live_order_product_options"

	// tickets
	TablePayments       string = "payments"
	TableTickets        string = "tickets"
	TableTicketLines    string = "ticket_lines"
	TableReturns        string = "returns"
	TableReturnLines    string = "return_lines"
	TableClosureReports string = "closure_reports"

	//payments
	TablePaymentMethods string = "payment_methods"
	TableTaxRates       string = "tax_rates"
)
