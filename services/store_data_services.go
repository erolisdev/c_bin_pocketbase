package services

import (
	"c_bin_pocketbase/constants"
	models "c_bin_pocketbase/models/store_models"
	"encoding/json"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func GetStoreData(app core.App) (*models.StoreData, error) {
	categories, err := GetStoreCategories(app)
	if err != nil {
		return nil, err
	}

	products, err := GetStoreProducts(app)
	if err != nil {
		return nil, err
	}

	storeData := models.StoreData{
		Categories: categories,
		Products:   products,
	}

	return &storeData, nil
}

func GetStoreCategories(app core.App) ([]models.StoreCategory, error) {

	var categories []models.StoreCategory
	var storeCategoryRecords []core.Record

	err := app.RecordQuery(constants.TableStoreCategories).
		Where(dbx.NewExp("status = {:status}", dbx.Params{"status": 1})).
		OrderBy("sort_order ASC").
		All(&storeCategoryRecords)

	if err != nil {
		return nil, err
	}

	for _, rec := range storeCategoryRecords {

		category := models.StoreCategory{
			CategoryID:      rec.GetInt("category_id"),
			Column:          rec.GetInt("column"),
			SortOrder:       rec.GetInt("sort_order"),
			Image:           rec.GetString("image_url"),
			ShowInSuggested: rec.GetInt("show_in_suggested"),
			Name:            rec.GetString("name"),
			Descriptions:    rec.Get("description"),
		}
		categories = append(categories, category)
	}

	return categories, nil

}

func GetStoreProducts(app core.App) ([]models.StoreProduct, error) {

	records := []*core.Record{}
	var storeProducts []models.StoreProduct

	err := app.RecordQuery(constants.TableStoreProducts).
		All(&records)

	if err != nil {
		return nil, apis.NewBadRequestError("Order verileri alınamadı "+err.Error(), err)
	}

	errs := app.ExpandRecords(records, []string{
		"tax_rate",
	}, nil)

	if len(errs) > 0 {
		return nil, apis.NewBadRequestError("Order verileri alınamadı: %v", errs)
	}

	var retrivedStoreOptions = make(map[string]models.StoreOption)

	for _, rec := range records {
		storeProduct := models.StoreProduct{
			ProductID:       rec.GetInt("product_id"),
			CategoryID:      rec.GetInt("category_id"),
			Image:           rec.GetString("image_url"),
			Shipping:        rec.GetInt("shipping"),
			Price:           rec.GetString("price_ttc"),
			PrinterID:       rec.GetInt("printer_id"),
			LblPrinterID:    rec.GetInt("lbl_printer_id"),
			SortOrder:       rec.GetInt("sort_order"),
			ShowInSuggested: rec.GetBool("show_in_suggested"),
			Descriptions:    rec.GetRaw("description"),
		}

		taxRate := rec.ExpandedOne("tax_rate")
		storeProduct.TaxRate = taxRate.GetString("rate")

		optionsRaw := rec.GetRaw("product_option_ids").(types.JSONRaw)

		var options []map[string]any
		err := json.Unmarshal(optionsRaw, &options)
		if err != nil {
			fmt.Println("Error unmarshaling options ids from product tables:", err)
			continue
		}

		// fmt.Println("Type of options:", reflect.TypeOf(options), "len: ", len(options))

		for _, option := range options {

			optionID := ToStringSimple(option["option"])
			if optionID == "" {
				continue
			}

			if _, exists := retrivedStoreOptions[optionID]; !exists {

				var storeOption models.StoreOption
				var storeOptionRecord core.Record

				optErr := app.RecordQuery(constants.TableStoreOptions).
					Where(dbx.NewExp("id = {:id}", dbx.Params{"id": optionID})).
					One(&storeOptionRecord)

				if optErr != nil {
					return nil, fmt.Errorf("Option not found: %s", optionID)
				}

				storeOption = models.StoreOption{
					OptionID:        storeOptionRecord.GetInt("id"),
					Required:        storeOptionRecord.GetInt("required"),
					FreeOptionCount: storeOptionRecord.GetInt("free_option_count"),
					MaxOptionCount:  storeOptionRecord.GetInt("max_option_count"),
					Type:            storeOptionRecord.GetString("type"),
					OptionGroup:     storeOptionRecord.GetInt("option_group"),
					Title:           storeOptionRecord.GetInt("title"),
					Descriptions:    storeOptionRecord.GetRaw("descriptions"),
				}

				// product option sort order
				sortOrder, err := parseToInt(option["sort_order"])
				if err == nil {
					storeOption.PoSortOrder = sortOrder
				} else {
					fmt.Println("Error sort order parseToInt: ", err.Error())
				}

				// order option values
				var optionValues []models.StoreOptionValue
				var optionValueRecords []core.Record

				err = app.RecordQuery(constants.TableStoreOptionValues).
					Where(dbx.NewExp("option = {:option}", dbx.Params{"option": optionID})).
					OrderBy("sort_order ASC").
					All(&optionValueRecords)

				if err != nil {
					return nil, err
				}

				for _, rec := range optionValueRecords {
					optionValue := models.StoreOptionValue{
						OptionID:        rec.GetInt("option_id"),
						OptionValueID:   rec.GetInt("option_value_id"),
						Image:           rec.GetString("image_url"),
						SortOrder:       rec.GetInt("sort_order"),
						PriceTTC:        rec.GetString("price_ttc"),
						PriceHT:         rec.GetString("price_ht"),
						Descriptions:    rec.GetRaw("descriptions"),
						ProductID:       &storeProduct.ProductID,
						PriceStatus:     rec.GetString("price_status"),
						Reset:           rec.GetInt("reset"),
						Grup:            rec.GetString("grup"),
						RelatedOptionID: rec.GetRaw("related_option_id"),
					}
					optionValues = append(optionValues, optionValue)
				}

				// fmt.Println("optionValueRecords: ", *optionValueRecords[0].Descriptions)

				storeOption.Values = &optionValues
				storeProduct.Options = append(storeProduct.Options, storeOption)

				retrivedStoreOptions[optionID] = storeOption

			} else {
				storeProduct.Options = append(storeProduct.Options, retrivedStoreOptions[optionID])
			}

		}

		storeProducts = append(storeProducts, storeProduct)

	}

	return storeProducts, nil
}
