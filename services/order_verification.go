package services

// import (
// 	"c_bin_pocketbase/models"
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/pocketbase/dbx"
// 	"github.com/pocketbase/pocketbase"
// 	"github.com/pocketbase/pocketbase/core"
// )

// // verifyOrder verifies the order against the database
// func VerifyOrder(app *pocketbase.PocketBase, order models.OrderModel) error {
// 	var calculatedOrderTotal float64  // Total based on prices sent by the client
// 	var calculatedRecordTotal float64 // Total based on prices from the DB and applied rules

// 	allProductIDs := make(map[int64]struct{})
// 	allOptionValueIDs := make(map[int64]struct{})

// 	// First pass: collect all unique product and option value IDs
// 	for _, product := range order.OrderData.OrderProducts {
// 		allProductIDs[*product.ProductID] = struct{}{}
// 		for _, opt := range product.OptionValues {
// 			allOptionValueIDs[*opt.OptionValueID] = struct{}{}
// 		}
// 	}

// 	productIDsSlice := make([]int64, 0, len(allProductIDs))
// 	for id := range allProductIDs {
// 		productIDsSlice = append(productIDsSlice, id)
// 	}
// 	optionValueIDsSlice := make([]int64, 0, len(allOptionValueIDs))
// 	for id := range allOptionValueIDs {
// 		optionValueIDsSlice = append(optionValueIDsSlice, id)
// 	}

// 	// --- Fetch Option Values and their parent Option rules from DB ---
// 	dbOptionValueRecords := []*core.Record{}
// 	// We need:
// 	// From option_values: id (record id), option_id, option_value_id, price_ttc
// 	// From options: type, free_count, max_count, name (option group name)
// 	err := app.RecordQuery("option_values").
// 		AndWhere(dbx.In("option_value_id", optionValueIDsSlice)). // Use variadic version of In
// 		InnerJoin("options", dbx.NewExp("option_values.option_id = options.id")).
// 		Select(
// 			"option_values.id",
// 			"option_values.option_id",
// 			"option_values.option_value_id",
// 			"option_values.price_ttc", // This is the price of the specific option_value
// 			"options.name as option_group_name",
// 			"options.type as option_type",
// 			"options.free_count as option_free_count",
// 			"options.max_count as option_max_count",
// 		).
// 		All(&dbOptionValueRecords)

// 	if err != nil {
// 		return fmt.Errorf("failed to fetch option values: %w", err)
// 	}

// 	// Map for quick lookup of DB data by OptionValueID
// 	dbOptionValuesMap := make(map[int64]models.DBDataOptionValue)
// 	// Map to store rules per OptionID (parent option)
// 	// Key: OptionID (e.g., "color_options_id"), Value: rules for "color"
// 	optionRulesMap := make(map[int64]struct {
// 		Type      string
// 		FreeCount int
// 		MaxCount  int
// 		Name      string
// 	})

// 	for _, rec := range dbOptionValueRecords {
// 		// ovID := rec.GetString("option_value_id")
// 		// optID := rec.GetString("option_id") // Parent option ID
// 		ovID := int64(rec.GetInt("option_value_id"))
// 		optID := int64(rec.GetInt("option_id")) // Parent option ID

// 		dbOptionValuesMap[ovID] = models.DBDataOptionValue{
// 			ID:            rec.Id,
// 			OptionID:      optID,
// 			OptionValueID: ovID,
// 			PriceTTC:      rec.GetFloat("price_ttc"),
// 			// Option group specific rules will be populated from optionRulesMap
// 		}

// 		if _, exists := optionRulesMap[optID]; !exists {
// 			optionRulesMap[optID] = struct {
// 				Type      string
// 				FreeCount int
// 				MaxCount  int
// 				Name      string
// 			}{
// 				Name:      rec.GetString("option_group_name"),
// 				Type:      rec.GetString("option_type"),
// 				FreeCount: rec.GetInt("option_free_count"),
// 				MaxCount:  rec.GetInt("option_max_count"),
// 			}
// 		}
// 	}

// 	// --- Fetch Product prices from DB ---
// 	dbProductRecords := []*core.Record{}
// 	productErr := app.RecordQuery("products").
// 		AndWhere(dbx.In("product_id", productIDsSlice)). // Use variadic version of In
// 		Select("product_id", "price_ttc").
// 		All(&dbProductRecords)

// 	if productErr != nil {
// 		return fmt.Errorf("failed to fetch products: %w", productErr)
// 	}

// 	dbProductPricesMap := make(map[int64]float64)
// 	for _, rec := range dbProductRecords {
// 		// This is the product ID from the products table
// 		dbProductPricesMap[int64(rec.GetInt("product_id"))] = rec.GetFloat("price_ttc")
// 	}

// 	// --- Second pass: Calculate totals and verify ---
// 	for _, orderProduct := range order.OrderData.OrderProducts {
// 		// if orderProduct.ProductID == nil || *orderProduct.ProductID <= 0 {
// 		// 	return fmt.Errorf("product ID is missing or invalid for product '%s'", orderProduct.Name)
// 		// }

// 		// Get product price from DB
// 		dbProductPrice, ok := dbProductPricesMap[*orderProduct.ProductID]
// 		if !ok {
// 			return fmt.Errorf("product with ID '%s' not found in database", orderProduct.ProductID)
// 		}

// 		// Accumulate prices for this product based on client's submitted prices
// 		// clientSideProductPriceTotal := orderProduct.Price // Price of one unit of the product
// 		clientSideProductPriceTotal, err := strconv.ParseFloat(*orderProduct.Price, 64)
// 		if err != nil {
// 			return fmt.Errorf("invalid product price for '%s': %v", orderProduct.Name, err)
// 		}

// 		clientSideOptionsPriceForOneProduct := 0.0

// 		// Accumulate prices for this product based on DB prices and rules
// 		dbSideProductPriceTotal := dbProductPrice // Price of one unit of the product from DB
// 		dbSideOptionsPriceForOneProduct := 0.0

// 		// Group option values from the order by their parent OptionID
// 		groupedOrderOptionValues := make(map[int64][]models.OrderOptionValue)
// 		for _, ov := range orderProduct.OptionValues {
// 			groupedOrderOptionValues[*ov.OptionID] = append(groupedOrderOptionValues[*ov.OptionID], ov)
// 		}

// 		// Process options for the current product
// 		for optionID, orderOVsInGroup := range groupedOrderOptionValues {
// 			rules, rulesExist := optionRulesMap[optionID]
// 			if !rulesExist {
// 				return fmt.Errorf("rules for option group ID '%s' not found for product '%s'", optionID, orderProduct.ProductID)
// 			}

// 			// MaxCount check: refers to the number of *distinct* option_values selected for this option group
// 			if rules.MaxCount > 0 && len(orderOVsInGroup) > rules.MaxCount {
// 				return fmt.Errorf(
// 					"product '%s': option group '%s' (ID: %s) allows max %d selections, but %d were made",
// 					orderProduct.Name, rules.Name, optionID, rules.MaxCount, len(orderOVsInGroup),
// 				)
// 			}

// 			// For "general" type options, track how many free items have been "used up" for this OptionID group
// 			freeItemsUsedForThisOptionGroup := 0

// 			for _, orderOV := range orderOVsInGroup { // orderOV is from the client's order
// 				dbOVData, dbOVExists := dbOptionValuesMap[*orderOV.OptionValueID]
// 				if !dbOVExists {
// 					return fmt.Errorf("option value with ID '%s' (for product '%s', option group '%s') not found in database",
// 						orderOV.OptionValueID, orderProduct.Name, rules.Name)
// 				}

// 				var orderValueQuantity float64
// 				if orderOV.Quantity == nil || *orderOV.Quantity <= 0 {
// 					orderValueQuantity = 1.0
// 				} else {
// 					orderValueQuantity = *orderOV.Quantity
// 				}

// 				orderOvPrice, err := strconv.ParseFloat(*orderOV.Price, 64)
// 				if err != nil {
// 					orderOvPrice = 0.0 // If parsing fails, treat it as zero
// 				}

// 				// Add client-submitted option price to the client-side total for one product unit
// 				// clientSideOptionsPriceForOneProduct += orderOV.Price * float64(orderOV.Quantity)
// 				clientSideOptionsPriceForOneProduct += orderOvPrice * orderValueQuantity

// 				// Calculate DB-side price for this option value, considering quantity and free counts
// 				for i := 0; i < int(orderValueQuantity); i++ { // Iterate for each unit of this option value
// 					isChargeable := true
// 					if rules.Type == "general" && rules.FreeCount > 0 {
// 						if freeItemsUsedForThisOptionGroup < rules.FreeCount {
// 							isChargeable = false
// 							freeItemsUsedForThisOptionGroup++
// 						}
// 					} else if rules.Type == "special" {
// 						// Special types are always chargeable, free_count is ignored
// 						isChargeable = true
// 					}
// 					// If rules.Type == "general" and rules.FreeCount <= 0, it's also always chargeable

// 					if isChargeable {
// 						dbSideOptionsPriceForOneProduct += dbOVData.PriceTTC // PriceTTC is per unit of option value
// 					}
// 				}
// 			}
// 		}

// 		orderProductQuentity := 1.0
// 		if orderProduct.Quantity != nil && *orderProduct.Quantity > 0 {
// 			orderProductQuentity = float64(*orderProduct.Quantity)
// 		}

// 		// Add to grand totals, multiplying by product quantity
// 		calculatedOrderTotal += (clientSideProductPriceTotal + clientSideOptionsPriceForOneProduct) * orderProductQuentity
// 		calculatedRecordTotal += (dbSideProductPriceTotal + dbSideOptionsPriceForOneProduct) * orderProductQuentity
// 	}

// 	orderInfoTotal, err := strconv.ParseFloat(*order.OrderData.Total, 64)
// 	if err != nil {
// 		return fmt.Errorf("invalid order total price: %v", err)
// 	}

// 	// Final Verifications
// 	// 1. Compare calculated total (based on DB prices) with the total price in the order
// 	if !areFloatsEqual(calculatedRecordTotal, orderInfoTotal, 0.01) {
// 		return fmt.Errorf("price mismatch: calculated DB total %.2f, order info total %.2f",
// 			calculatedRecordTotal, orderInfoTotal)
// 	}

// 	// 2. Compare client's internal calculation (order total vs sum of its parts)
// 	//    This check is good for sanity but calculatedRecordTotal vs order.OrderInfo.TotalPrice is the primary one.
// 	if !areFloatsEqual(calculatedRecordTotal, calculatedOrderTotal, 0.01) {
// 		// This might indicate that the client sent inconsistent pricing data for its own items,
// 		// or our interpretation of clientSideOptionsPriceForOneProduct is slightly off due to float math.
// 		// The more important check is that calculatedRecordTotal matches order.OrderInfo.TotalPrice.
// 		// If they don't match, it means the client's "total" doesn't match what they should pay based on DB prices.
// 		// If calculatedRecordTotal == order.OrderInfo.TotalPrice, but calculatedRecordTotal != calculatedOrderTotal,
// 		// it means the client's order.OrderInfo.TotalPrice is correct according to DB, but their itemized prices might be off.
// 		// For strictness, we can keep this check.
// 		log.Printf("Debug: calculatedRecordTotal: %.5f, calculatedOrderTotal: %.5f", calculatedRecordTotal, calculatedOrderTotal)
// 		return fmt.Errorf("internal price consistency mismatch: calculated DB total %.2f, calculated client-side total %.2f. This might indicate client sent item prices that don't sum up to their own OrderInfo.TotalPrice, or a calculation discrepancy.",
// 			calculatedRecordTotal, calculatedOrderTotal)
// 	}

// 	// Order date format check can be re-enabled if needed
// 	// if _, err := time.Parse("2006-01-02T15:04:05Z07:00", order.OrderInfo.OrderDate); err != nil { // Example: RFC3339
// 	// 	return fmt.Errorf("invalid order date format: %v", err)
// 	// }

// 	return nil
// }
