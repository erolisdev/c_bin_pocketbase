package services

// import (
// 	"c_bin_pocketbase/models"
// 	"fmt"

// 	"github.com/labstack/echo"
// 	"github.com/pocketbase/dbx"
// 	"github.com/pocketbase/pocketbase/core"
// 	"github.com/spf13/cast"
// )

// func FetchAndStructureOrderData(app core.App, clientOrder *models.OrderModel, event *core.RequestEvent) (*models.OrderModel, error) {
// 	if clientOrder == nil || len(clientOrder.OrderProducts) == 0 {
// 		return nil, event.BadRequestError("Order has no products", "error") // veya e.BadRequestError
// 	}

// 	finalOrderModel := &models.OrderModel{
// 		OrderProducts: []models.OrderProduct{},
// 		// ... Diğer OrderModel alanlarını doldurun (müşteri ID, toplam tutar vb.)
// 	}

// 	// Önbellekler
// 	groupedProducts := make(map[int]*core.Record)
// 	groupedOptions := make(map[int]*core.Record)      // Ana seçenekler için (örn: Renk)
// 	groupedOptionValues := make(map[int]*core.Record) // Seçenek değerleri için (örn: Kırmızı)

// 	for _, prdInput := range clientOrder.OrderProducts {
// 		productIDStr := prdInput.ProductID
// 		if productIDStr < 0 {
// 			return nil, echo.NewHTTPError(400, "Invalid product data: ProductID is missing")
// 		}

// 		var productRecord *core.Record
// 		var err error

// 		// 1. Ürünü Getir/Önbellekten Al
// 		if cachedRecord, found := groupedProducts[productIDStr]; found {
// 			productRecord = cachedRecord
// 		} else {
// 			productRecord, err = app.Dao().FindFirstRecordByFilter(
// 				productsCollection,
// 				"id = {:id}", // PocketBase ID'si ile filtreleme
// 				dbx.Params{"id": productIDStr},
// 			)
// 			if err != nil {
// 				return nil, echo.NewHTTPError(404, fmt.Sprintf("Product not found (ID: %s): %s", productIDStr, err.Error()))
// 			}
// 			if productRecord == nil { // Ekstra kontrol, FindFirstRecordByFilter nil dönebilir
// 				return nil, echo.NewHTTPError(404, fmt.Sprintf("Product not found (ID: %s), record is nil", productIDStr))
// 			}
// 			groupedProducts[productIDStr] = productRecord
// 		}

// 		// models.OrderProduct oluştur ve DB'den gelen verilerle doldur
// 		orderProduct := models.OrderProduct{
// 			ProductID: prdInput.ProductID, // İstemciden gelen ID'yi koru
// 			Name:      types.StringPointer(productRecord.GetString(productNameField)),
// 			Quantity:  prdInput.Quantity,                                                                 // İstemciden gelen miktarı kullan
// 			UpriceTTC: types.Float64Pointer(cast.ToFloat64(productRecord.Get(productUnitPriceTTCField))), // Örnek
// 			UpriceHT:  types.Float64Pointer(cast.ToFloat64(productRecord.Get(productUnitPriceHTField))),  // Örnek
// 			// Status, PrixRevt, TotalTVA, TotalHT, TotalTTC, TaxRate gibi alanları
// 			// ya productRecord'dan ya da hesaplamalarla doldurun.
// 			// Şimdilik bazılarını boş bırakıyorum veya varsayılan değerlerle.
// 		}

// 		// 2. Ürün Seçeneklerini İşle
// 		if prdInput.Options != nil && len(*prdInput.Options) > 0 {
// 			processedOptions := []models.OrderOption{}
// 			for _, optInput := range *prdInput.Options {
// 				optionIDStr := safeStringPtrValue(optInput.OptionID)
// 				if optionIDStr == "" {
// 					return nil, echo.NewHTTPError(400, "Invalid option data: OptionID is missing for product "+productIDStr)
// 				}

// 				var optionRecord *core.Record // Bu ana seçeneği temsil eder (örn: Renk)

// 				// 2a. Ana Seçeneği Getir/Önbellekten Al
// 				if cachedRecord, found := groupedOptions[optionIDStr]; found {
// 					optionRecord = cachedRecord
// 				} else {
// 					optionRecord, err = app.Dao().FindFirstRecordByFilter(
// 						optionsCollection,
// 						"id = {:id}",
// 						dbx.Params{"id": optionIDStr},
// 					)
// 					if err != nil {
// 						return nil, echo.NewHTTPError(404, fmt.Sprintf("Option not found (ID: %s): %s", optionIDStr, err.Error()))
// 					}
// 					if optionRecord == nil {
// 						return nil, echo.NewHTTPError(404, fmt.Sprintf("Option not found (ID: %s), record is nil", optionIDStr))
// 					}
// 					groupedOptions[optionIDStr] = optionRecord
// 				}

// 				orderOption := models.OrderOption{
// 					OptionID:     optInput.OptionID,  // İstemciden gelen ID'yi koru
// 					ProductID:    prdInput.ProductID, // Üst ürün ID'si
// 					Name:         types.StringPointer(optionRecord.GetString(optionNameField)),
// 					OptionValues: []models.OrderOptionValue{},
// 					// Diğer OrderOption alanlarını (Unit, Quantity vs.)
// 					// ya optionRecord'dan ya da hesaplamalarla doldurun.
// 					// Genellikle bu alanlar OptionValue seviyesinde olur.
// 				}

// 				// 2b. Seçenek Değerlerini İşle
// 				if len(optInput.OptionValues) > 0 {
// 					for _, optValueInput := range optInput.OptionValues {
// 						optionValueIDStr := safeStringPtrValue(optValueInput.OptionValueID)
// 						if optionValueIDStr == "" {
// 							return nil, echo.NewHTTPError(400, fmt.Sprintf("Invalid option value data: OptionValueID is missing for option %s, product %s", optionIDStr, productIDStr))
// 						}

// 						var optionValueRecord *core.Record

// 						// Seçenek Değerini Getir/Önbellekten Al
// 						if cachedRecord, found := groupedOptionValues[optionValueIDStr]; found {
// 							optionValueRecord = cachedRecord
// 						} else {
// 							optionValueRecord, err = app.Dao().FindFirstRecordByFilter(
// 								optionValuesCollection,
// 								"id = {:id}",
// 								dbx.Params{"id": optionValueIDStr},
// 							)
// 							if err != nil {
// 								return nil, echo.NewHTTPError(404, fmt.Sprintf("Option value not found (ID: %s): %s", optionValueIDStr, err.Error()))
// 							}
// 							if optionValueRecord == nil {
// 								return nil, echo.NewHTTPError(404, fmt.Sprintf("Option value not found (ID: %s), record is nil", optionValueIDStr))
// 							}
// 							groupedOptionValues[optionValueIDStr] = optionValueRecord
// 						}

// 						orderOptionValue := models.OrderOptionValue{
// 							OptionValueID: optValueInput.OptionValueID, // İstemciden gelen ID'yi koru
// 							OptionID:      optInput.OptionID,           // Ana seçenek ID'si
// 							ProductID:     prdInput.ProductID,          // Üst ürün ID'si
// 							Name:          types.StringPointer(optionValueRecord.GetString(optionValueNameField)),
// 							Quantity:      optValueInput.Quantity,                                                                    // İstemciden gelen miktarı kullan
// 							UpriceTTC:     types.Float64Pointer(cast.ToFloat64(optionValueRecord.Get(optionValueUnitPriceTTCField))), // Örnek
// 							// Diğer OrderOptionValue alanlarını (Unit, UpriceHTC, PrixRevt, TotalTVA, TotalHT, TotalTTC, TaxRate)
// 							// ya optionValueRecord'dan ya da hesaplamalarla doldurun.
// 						}
// 						// Eğer seçenek değeri fiyatları etkiliyorsa ve hesaplanması gerekiyorsa, burada yapılabilir.

// 						orderOption.OptionValues = append(orderOption.OptionValues, orderOptionValue)
// 					}
// 				}
// 				processedOptions = append(processedOptions, orderOption)
// 			}
// 			orderProduct.Options = &processedOptions
// 		}
// 		finalOrderModel.OrderProducts = append(finalOrderModel.OrderProducts, orderProduct)
// 	}

// 	// Burada siparişin genel toplamlarını hesaplayabilirsiniz (finalOrderModel.GrandTotalTTC vb.)
// 	// ...

// 	return finalOrderModel, nil
// }
