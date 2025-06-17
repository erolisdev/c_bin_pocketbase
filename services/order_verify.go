package services

import (
	"c_bin_pocketbase/models"
	"fmt"
	"log"
	"math"

	"c_bin_pocketbase/constants"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// verifyOrder verifies the order against the database
func VerifyOrder(app *pocketbase.PocketBase, orderInput models.OrderModel) (*models.OrderData, error) {
	if orderInput.OrderData == nil {
		return nil, fmt.Errorf("order data (items) is missing")
	}

	orderData := orderInput.OrderData

	var calculatedOrderTotal float64  // Total based on prices sent by the client
	var calculatedRecordTotal float64 // Total based on prices from the DB and applied rules

	allProductIDsInOrder := make(map[int64]struct{})     // Key: string representation of ProductID
	allOptionValueIDsInOrder := make(map[int64]struct{}) // Key: string representation of ProductOptionValueID

	// First pass: collect all unique product and option value IDs from the input order
	for _, orderProduct := range orderData.OrderProducts {
		if orderProduct.ProductID == nil {
			return nil, fmt.Errorf("product is missing ProductID")
		}
		allProductIDsInOrder[*orderProduct.ProductID] = struct{}{}

		for _, orderOV := range orderProduct.OptionValues {
			if orderOV.OptionValueID == nil { // Bu ID'yi option_values.option_value_id'ye map ettiğimizi varsayıyoruz
				return nil, fmt.Errorf("option value for product '%s' is missing ProductOptionValueID", derefString(orderProduct.Name, safeStringFromInt64Ptr(orderProduct.ProductID)))
			}
			allOptionValueIDsInOrder[*orderOV.OptionValueID] = struct{}{}
		}
	}

	productIDsSlice := make([]int64, 0, len(allProductIDsInOrder))
	for id := range allProductIDsInOrder {
		productIDsSlice = append(productIDsSlice, id)
	}
	optionValueIDsSlice := make([]int64, 0, len(allOptionValueIDsInOrder))
	for id := range allOptionValueIDsInOrder {
		optionValueIDsSlice = append(optionValueIDsSlice, id)
	}

	// --- Fetch Option Values and their parent Option rules from DB ---
	dbOptionValueRecords := []*core.Record{}
	if len(optionValueIDsSlice) > 0 { // Sadece ID varsa sorgula
		err := app.RecordQuery(constants.TableStoreOptionValues).
			Where(dbx.In("option_value_id", convertToInterfaceSlice(optionValueIDsSlice)...)).
			All(&dbOptionValueRecords)

		errs := app.ExpandRecords(dbOptionValueRecords, []string{
			"option", "tax_rate", // expand relation
		}, nil)

		if len(errs) > 0 {
			return nil, apis.NewBadRequestError("failed to expand option values: %w", errs)
		}

		// for _, rec := range dbOptionValueRecords {
		// 	fmt.Println("==============================")
		// 	fmt.Println("Option Value Expanded One:", rec.PublicExport())
		// 	fmt.Println("Option Value Expanded All:", rec.ExpandedAll("option"))
		// 	for _, rec2 := range rec.ExpandedAll("option") {
		// 		fmt.Println("Option Expanded OPTION:", rec2.PublicExport())
		// 	}
		// 	for _, rec3 := range rec.ExpandedAll("tax_rate") {
		// 		fmt.Println("Option Expanded TAX RATE:", rec3.PublicExport())
		// 	}

		// 	fmt.Println("================ db finish  ==============")
		// }

		if err != nil {
			return nil, fmt.Errorf("failed to fetch option values: %w", err)
		}
	}

	dbOptionValuesMap := make(map[int64]models.DBDataOptionValue) // Key: string(ProductOptionValueID)
	optionRulesMap := make(map[int64]struct {                     // Key: string(OptionID from DB)
		OptionType string
		FreeCount  int
		MaxCount   int
		Name       string
		// PriceStatus strings
	})

	for _, rec := range dbOptionValueRecords {
		// rec.GetString("option_value_id") -> PB'deki "option_values" tablosundaki "option_value_id" alanı
		// Bu, bizim OrderOptionValue.ProductOptionValueID'mize karşılık gelir (string'e çevrilmiş hali).
		dbOptionValueID := int64(rec.GetInt("option_value_id")) // This is the key from DB, e.g., "red_option"
		// dbOptionID := int64(rec.GetInt("option_id"))

		// we need to add option_id on table option_values
		optionRec := rec.ExpandedOne("option")
		dbOptionID := int64(optionRec.GetInt("option_id"))

		// fmt.Println("==============================")
		// fmt.Println("1 -dbOptionValueRecords Public Export:", dbOptionID, rec.PublicExport())
		// fmt.Println("==============================")

		if _, exists := optionRulesMap[dbOptionID]; !exists {

			// fmt.Println("==============================")
			// fmt.Println("2 -dbOptionValueRecords Public Export:", rec.PublicExport())
			// fmt.Println("==============================")
			//expand the option record to get its rules
			// optionRec := rec.ExpandedOne("option") // we need to add option_id on table option_values
			optionType := optionRec.GetString("type")
			maxCount := optionRec.GetInt("max_count")

			if optionType == models.Radio {
				maxCount = 1
			}

			optionRulesMap[dbOptionID] = struct {
				OptionType string
				FreeCount  int
				MaxCount   int
				Name       string
				// PriceStatus string
			}{
				Name:       optionRec.GetString("name"), // option name
				OptionType: optionRec.GetString("type"),
				// PriceStatus: rec.GetString("price_status"), // Assuming "price_status" is used to determine the type
				FreeCount: optionRec.GetInt("free_option_count"),
				MaxCount:  maxCount,
			}

			// log.Println("======================== =============================")
			// log.Println("Option Expanded All:", rec.ExpandedAll("option"))
			// log.Println("======================== =============================")
			// log.Println("Option Expanded One:", optionRec.PublicExport())
			// log.Println("======================== ", optionRulesMap[dbOptionID].Name, "============================")

			// log.Println("Option rules added for OptionID:", dbOptionID,
			// 	// "Type:", optionRulesMap[dbOptionID].Name,
			// 	"FreeCount:", optionRulesMap[dbOptionID].FreeCount,
			// 	"MaxCount:", optionRulesMap[dbOptionID].MaxCount,
			// 	"Name:", optionRulesMap[dbOptionID].Name,
			// 	"OptionID:", dbOptionID,
			// 	"OptionValueID:", dbOptionValueID,
			// )

			// log.Println("======================== =============================")
			// Note: We assume that the name is the same for all records with the same OptionID
		}

		dbOptionValuesMap[dbOptionValueID] = models.DBDataOptionValue{
			// optionvalue data
			OptionID:      dbOptionID,
			OptionValueID: dbOptionValueID,
			PriceTTC:      rec.GetFloat("price_ttc"),
			Name:          rec.GetString("name"),
			PriceStatus:   rec.GetString("price_status"),
			// parent option data
			OptionType: optionRulesMap[dbOptionID].OptionType,
			FreeCount:  optionRulesMap[dbOptionID].FreeCount,
			MaxCount:   optionRulesMap[dbOptionID].MaxCount,
		}

		// Rules are per parent option (OptionID)

	}

	// --- Fetch Product prices from DB ---
	// products.product_id (string): Benzersiz ürün ID'si (bizim modelde ProductID)
	// products.price_ttc (number): Ürünün fiyatı
	dbProductRecords := []*core.Record{}
	if len(productIDsSlice) > 0 {
		productErr := app.RecordQuery(constants.TableStoreProducts).
			AndWhere(dbx.In("product_id", convertToInterfaceSlice(productIDsSlice)...)).
			Select("product_id", "price_ttc", "name", "tax_rate").
			All(&dbProductRecords)

		if productErr != nil {
			return nil, fmt.Errorf("failed to fetch products: %w", productErr)
		}

		for _, rec := range dbProductRecords {
			log.Println("======================== =============================")
			log.Println("Prd rec: ", rec.PublicExport())
			// log.Println("Prd Expanded All:", rec.ExpandedOne("tax_rate").PublicExport())
			log.Println("======================== =============================")
		}
	}

	// dbProductPricesMap := make(map[string]float64) // Key: string(ProductID)
	dbProductPricesMap := make(map[string]*core.Record) // Key: string(ProductID)
	for _, rec := range dbProductRecords {
		// dbProductPricesMap[rec.GetString("product_id")] = rec.GetFloat("price_ttc")
		dbProductPricesMap[rec.GetString("product_id")] = rec
	}

	// --- Second pass: Calculate totals and verify ---
	for pIndex, orderProduct := range orderData.OrderProducts {
		orderProductIDStr := safeStringFromInt64Ptr(orderProduct.ProductID)

		dbProductRecord, ok := dbProductPricesMap[orderProductIDStr]
		if !ok {
			return nil, fmt.Errorf("product with ID '%s' (name: %s) not found in database", orderProductIDStr, derefString(orderProduct.Name, "N/A"))
		}
		// dbProductPrice, ok := dbProductPricesMap[orderProductIDStr]
		// if !ok {
		// 	return nil, fmt.Errorf("product with ID '%s' (name: %s) not found in database", orderProductIDStr, derefString(orderProduct.Name, "N/A"))
		// }

		dbProductPrice := dbProductRecord.GetFloat("price_ttc")
		// dbProductPriceTaxRate := dbProductRecord.GetFloat("tax_rate")

		// for real product name
		dbProductName := dbProductRecord.GetString("name")
		orderData.OrderProducts[pIndex].Name = &dbProductName

		dbProductTaxRateID := dbProductRecord.GetString("tax_rate")
		orderData.OrderProducts[pIndex].TaxRate = &dbProductTaxRateID

		clientProductPrice, err := parseFloat(orderProduct.Price, fmt.Sprintf("product '%s' price", orderProductIDStr))
		if err != nil {
			return nil, err
		}
		orderProductQuantity := derefInt64(orderProduct.Quantity, 1)
		if orderProductQuantity <= 0 {
			return nil, fmt.Errorf("product '%s' has invalid quantity: %d", orderProductIDStr, orderProductQuantity)
		}

		clientSideOptionsPriceForOneProductUnit := 0.0
		dbSideOptionsPriceForOneProductUnit := 0.0

		// Group option values from the order by their parent OptionID (from input model)
		// The OptionID in the input model (OrderOptionValue.OptionID) should map to the parent option group in DB.
		groupedOrderOptionValues := make(map[int64][]models.OrderOptionValue) // Key: string(OrderOptionValue.OptionID)
		for ovIndex, ov := range orderProduct.OptionValues {
			if ov.OptionID == nil {
				return nil, fmt.Errorf("option value '%s' for product '%s' is missing its parent OptionID",
					safeStringFromInt64Ptr(ov.ProductOptionValueID), orderProductIDStr)
			}
			inputOptionID := ov.OptionID
			groupedOrderOptionValues[*inputOptionID] = append(groupedOrderOptionValues[*inputOptionID], ov)

			// set real option name
			dbOVData, dbOVExists := dbOptionValuesMap[*ov.OptionValueID]

			if dbOVExists {
				orderData.OrderProducts[pIndex].OptionValues[ovIndex].Name = &dbOVData.Name
			}
		}

		for inputOptionID, orderOVsInGroup := range groupedOrderOptionValues {
			// inputOptionIDStr -> Bu, istemcinin gönderdiği OptionID. Bunun DB'deki option_id'ye karşılık geldiğini varsayıyoruz.
			rules, rulesExist := optionRulesMap[inputOptionID] // Use inputOptionIDStr as key if it maps directly to DB's option_id
			fmt.Println("Processing option group ID:", inputOptionID, "with name", rules.Name, "rules", rules, "max count", rules.MaxCount, "free count", rules.FreeCount, "type")

			if !rulesExist {
				var firstOVName string
				if len(orderOVsInGroup) > 0 {
					firstOVName = derefString(orderOVsInGroup[0].Name, safeStringFromInt64Ptr(orderOVsInGroup[0].OptionValueID))
				}
				return nil, fmt.Errorf("rules for option group ID '%s' (e.g., for option value '%s' on product '%s') not found. Ensure OptionID from input maps to a valid option group in DB.",
					ToStringSimple(inputOptionID), firstOVName, orderProductIDStr)
			}

			if rules.MaxCount > 0 && len(orderOVsInGroup) > rules.MaxCount {

				return nil, fmt.Errorf(
					"product '%s': option group '%s' (ID: %s) allows max %d distinct selections, but %d were made",
					derefString(orderProduct.Name, orderProductIDStr), rules.Name, ToStringSimple(inputOptionID), rules.MaxCount, len(orderOVsInGroup),
				)
			}

			freeItemsUsedForThisOptionGroup := 0

			for _, orderOV := range orderOVsInGroup {

				dbOVData, dbOVExists := dbOptionValuesMap[*orderOV.OptionValueID]
				if !dbOVExists {
					return nil, fmt.Errorf("option value with ID '%s' (name: %s, for product '%s', option group '%s') not found in database",
						safeStringFromInt64Ptr(orderOV.OptionValueID), derefString(orderOV.Name, "N/A"), orderProductIDStr, rules.Name)
				}
				// DB'den gelen OptionID ile kuralları aldığımız OptionID'nin eşleştiğinden emin olalım (tutarlılık için)
				if dbOVData.OptionID != inputOptionID {
					return nil, fmt.Errorf("data integrity issue: option value '%s' (DB parent OptionID: %s) is grouped under a different OptionID ('%s') from input for product '%s'. Check input data or DB setup.",
						safeStringFromInt64Ptr(orderOV.OptionValueID), ToStringSimple(dbOVData.OptionID), ToStringSimple(inputOptionID), orderProductIDStr)
				}

				clientOptionPrice, err := parseFloat(orderOV.Price, fmt.Sprintf("option '%s' price", safeStringFromInt64Ptr(orderOV.ProductOptionValueID)))
				if err != nil {
					return nil, err
				}
				orderOptionQuantity := derefFloat64(orderOV.Quantity, 1.0)
				if orderOptionQuantity <= 0 {
					return nil, fmt.Errorf("option '%s' for product '%s' has invalid quantity: %f", safeStringFromInt64Ptr(orderOV.ProductOptionValueID), orderProductIDStr, orderOptionQuantity)
				}

				// fmt.Println("Processing option value:", ToStringSimple(orderOV.Name), "with client price", clientOptionPrice, "and quantity", orderOptionQuantity, "db price", dbOVData.PriceTTC)

				clientSideOptionsPriceForOneProductUnit += clientOptionPrice * orderOptionQuantity

				if rules.OptionType == models.Radio {
					dbSideOptionsPriceForOneProductUnit += dbOVData.PriceTTC
					fmt.Println(rules.OptionType, dbOVData.Name, "price:", dbOVData.PriceTTC, "dbSideOptionsPriceForOneProductUnit", dbSideOptionsPriceForOneProductUnit)
				} else {
					// DB-side price calculation
					// orderOptionQuantity float64 olduğu için, tam sayı adetler üzerinden ücretsiz kontrolü

					for i := 0; i < int(math.Round(orderOptionQuantity)); i++ { // Round to nearest int for quantity
						isChargeable := true

						if dbOVData.PriceStatus == models.General && rules.FreeCount > 0 {
							if freeItemsUsedForThisOptionGroup < rules.FreeCount {
								isChargeable = false
								freeItemsUsedForThisOptionGroup++
							}
						} else if dbOVData.PriceStatus == models.Special { // "special" types are always chargeable per item
							isChargeable = true
						} else if dbOVData.PriceStatus == models.OneIsFree {

							if len(orderOVsInGroup) == 1 {
								isChargeable = false
							} else {
								//TODO
							}
						}

						// If rules.Type == "general" and rules.FreeCount <= 0, it's also always chargeable
						if isChargeable {
							dbSideOptionsPriceForOneProductUnit += dbOVData.PriceTTC // PriceTTC is per unit of option value
						}

						fmt.Println(rules.OptionType, dbOVData.Name, "price:", dbOVData.PriceTTC, "dbSideOptionsPriceForOneProductUnit", dbSideOptionsPriceForOneProductUnit)

					}
					fmt.Println("dbSideOptionsPriceForOneProductUnit", dbSideOptionsPriceForOneProductUnit)

				}

			}
		}

		calculatedOrderTotal += (clientProductPrice + clientSideOptionsPriceForOneProductUnit) * float64(orderProductQuantity)
		calculatedRecordTotal += (dbProductPrice + dbSideOptionsPriceForOneProductUnit) * float64(orderProductQuantity)
	}

	// Final Verifications
	orderTotalFromInput, err := parseFloat(orderData.Total, "order total")
	if err != nil {
		return nil, err
	}

	if !areFloatsEqual(calculatedRecordTotal, orderTotalFromInput, 0.01) {
		return nil, fmt.Errorf("price mismatch: calculated DB total %.2f, order info total %.2f (from input string '%s')",
			calculatedRecordTotal, orderTotalFromInput, derefString(orderData.Total, "N/A"))
	}

	// This check becomes less critical if the above one passes, but can catch internal client calculation errors.
	if !areFloatsEqual(calculatedRecordTotal, calculatedOrderTotal, 0.01) {
		log.Printf("Debug: calculatedRecordTotal: %.5f, calculatedOrderTotal: %.5f", calculatedRecordTotal, calculatedOrderTotal)
		return nil, fmt.Errorf("internal price consistency mismatch: calculated DB total %.2f, calculated client-side total %.2f. This might indicate client sent item prices that don't sum up to their own OrderInfo.TotalPrice, or a calculation discrepancy.",
			calculatedRecordTotal, calculatedOrderTotal)
	}

	return orderData, nil
}
