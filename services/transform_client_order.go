package services

// import (
// 	"c_bin_pocketbase/models"
// )

// // TransformOrderProductOptions, sipariş ürünlerini ve seçeneklerini yeniden yapılandırır.
// // Seçenekleri OptionID'ye göre gruplar ve seçeneği olmayan ürünleri de korur.
// func TransformOrderProductOptions(orderModel *models.OrderModel) *models.OrderModel {
// 	if orderModel == nil {
// 		return nil
// 	}

// 	processedProducts := []models.OrderProduct{}

// 	for _, prd := range orderModel.OrderProducts {
// 		// Temel ürün bilgilerini kopyala
// 		newProduct := models.OrderProduct{
// 			ProductID: prd.ProductID,
// 			Name:      prd.Name,
// 			Unit:      prd.Unit,
// 			Status:    prd.Status,
// 			Quantity:  prd.Quantity,
// 			UpriceTTC: prd.UpriceTTC,
// 			UpriceHT:  prd.UpriceHT,
// 			PrixRevt:  prd.PrixRevt,
// 			TotalTVA:  prd.TotalTVA,
// 			TotalHT:   prd.TotalHT,
// 			TotalTTC:  prd.TotalTTC,
// 			TaxRate:   prd.TaxRate,
// 			Options:   nil,
// 		}

// 		if prd.Options != nil && len(*prd.Options) > 0 {
// 			// OptionID'yi anahtar olarak kullanan bir map. Değer olarak *models.OrderOption tutar
// 			// Böylece map'teki option'ı doğrudan güncelleyebiliriz.
// 			groupedOptions := make(map[int64]*models.OrderOption)

// 			for _, optSource := range *prd.Options { // optSource, orijinal seçenek değerini temsil eder
// 				// Gerekli alanların nil olup olmadığını kontrol et (özellikle OptionID)
// 				if optSource.OptionID == nil {
// 					// Loglama yapılabilir: fmt.Println("Uyarı: OptionID nil olan bir seçenek atlandı.")
// 					continue // OptionID olmadan gruplama yapamayız
// 				}
// 				optionID := *optSource.OptionID

// 				// Yeni OrderOptionValue oluştur
// 				// Not: optSource'taki alanların pointer olup olmadığına göre safe* fonksiyonları kullanılmalı
// 				orderOptionValue := models.OrderOptionValue{
// 					ProductID:     prd.ProductID,
// 					OptionID:      optSource.OptionID,
// 					OptionValueID: optSource.OptionValueID,
// 					Unit:          optSource.Unit,
// 					Quantity:      optSource.Quantity,
// 					UpriceTTC:     optSource.UpriceTTC,
// 					UpriceHTC:     optSource.UpriceHTC,
// 					PrixRevt:      optSource.PrixRevt,
// 					TotalTVA:      optSource.TotalTVA,
// 					TotalHT:       optSource.TotalHT,
// 					TotalTTC:      optSource.TotalTTC,
// 					TaxRate:       optSource.TaxRate,
// 					Name:          optSource.Name,
// 				}

// 				// Map'te bu OptionID'ye sahip bir OrderOption var mı diye kontrol et
// 				if existingOrderOption, found := groupedOptions[optionID]; found {
// 					// Varsa, OptionValue'yu mevcut OrderOption'a ekle
// 					existingOrderOption.OptionValues = append(existingOrderOption.OptionValues, orderOptionValue)
// 				} else {
// 					// Yoksa, yeni bir OrderOption oluştur ve map'e ekle
// 					newOrderOption := models.OrderOption{
// 						ProductID:    prd.ProductID,
// 						OptionID:     optSource.OptionID,
// 						Name:         optSource.Name,
// 						OptionValues: []models.OrderOptionValue{orderOptionValue},
// 						// Unit:         optSource.Unit,
// 						// Quantity:     optSource.Quantity, // Bu genellikle OptionValue'da olur
// 						// UpriceTTC: optSource.UpriceTTC,
// 						// UpriceHTC: optSource.UpriceHTC,
// 						// PrixRevt:  optSource.PrixRevt,
// 						// TotalTVA:  optSource.TotalTVA,
// 						// TotalHT:   optSource.TotalHT,
// 						// TotalTTC:  optSource.TotalTTC,
// 						// TaxRate:   optSource.TaxRate,
// 					}
// 					groupedOptions[optionID] = &newOrderOption
// 				}
// 			}

// 			// Map'teki OrderOption'ları bir slice'a dönüştür
// 			if len(groupedOptions) > 0 {
// 				productOptionsSlice := make([]models.OrderOption, 0, len(groupedOptions))
// 				for _, orderOptionPtr := range groupedOptions {
// 					productOptionsSlice = append(productOptionsSlice, *orderOptionPtr)
// 				}
// 				newProduct.Options = &productOptionsSlice
// 			}
// 		}
// 		// Seçenekli veya seçeneksiz ürünü listeye ekle
// 		processedProducts = append(processedProducts, newProduct)
// 	}

// 	orderModel.OrderProducts = processedProducts
// 	if len(orderModel.OrderProducts) == 0 {
// 		// Eğer başlangıçta hiç ürün yoksa veya hepsi filtrelendiyse (ki burada filtreleme yok)
// 		// bu durumun nasıl ele alınacağına karar verilmeli.
// 		// Orijinal davranış: eğer sonuçta ürün yoksa nil dön.
// 		return nil
// 	}
// 	return orderModel
// }
