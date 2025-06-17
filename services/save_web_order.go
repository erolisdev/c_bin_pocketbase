package services

import (
	"c_bin_pocketbase/constants"
	"c_bin_pocketbase/models"
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func SaveOrder(app *pocketbase.PocketBase, orderData *models.OrderData) (*int, error) {

	liveOrderCollection, err := app.FindCollectionByNameOrId(constants.TableLiveOrders)
	if err != nil {
		return nil, err
	}

	liveOrderProductCollection, err := app.FindCollectionByNameOrId(constants.TableLiveOrderProducts)
	if err != nil {
		return nil, err
	}

	liveOrderProductOptionValueCollection, err := app.FindCollectionByNameOrId(constants.TableLiveOrderProductOptions)
	if err != nil {
		return nil, err
	}

	err = app.RunInTransaction(func(txApp core.App) error {

		if err != nil {
			return fmt.Errorf("live_orders save order number: '%s'", err.Error())
		}
		orderNumber, err := GetOrderNumber(txApp)
		orderData.OrderNumber = orderNumber

		fmt.Println("========= Save order number: ", *orderData.OrderNumber, ", ", orderNumber, "==========")

		// orderID, err := saveOrderInfo(txApp, liveOrderCollection, *orderData)
		orderID, orderNumber, err := saveOrderInfoWithRetry(txApp, liveOrderCollection, *orderData, 1)

		if err != nil {
			return fmt.Errorf("live_orders save error: '%s'", err.Error())
		}

		for _, product := range orderData.OrderProducts {
			productID, err := saveOrderProduct(txApp, liveOrderProductCollection, product, orderID)
			if err != nil {
				return fmt.Errorf("live_order_products save error: '%s'", err.Error())
			}

			for _, optionValue := range product.OptionValues {
				if err := saveOrderProductOptionValue(txApp, liveOrderProductOptionValueCollection, optionValue, orderID, productID); err != nil {
					return fmt.Errorf("live_order_product_option_value save error: '%s'", err.Error())
				}
			}
		}

		return nil
	})

	return orderData.OrderNumber, err

}

func saveOrderInfoWithRetry(txApp core.App, collection *core.Collection, orderData models.OrderData, tryCount int) (string, *int, error) {
	if tryCount > 10 {
		return "", nil, fmt.Errorf("max retry count reached")
	}

	fmt.Printf("Saving order info (try %d)\n", tryCount)

	orderID, err := saveOrderInfo(txApp, collection, orderData)
	if err != nil {
		if strings.Contains(err.Error(), "date: Value must be unique; order_number: Value must be unique.") {
			// Burada yeni bir order number üretmelisiniz.
			// Örnek:
			orderData.OrderNumber, err = GetOrderNumber(txApp)

			fmt.Printf("Retrying with new order number: %d (try %d)\n", orderData.OrderNumber, tryCount+1)
			return saveOrderInfoWithRetry(txApp, collection, orderData, tryCount+1)
		}
		return "", nil, err
	}

	return orderID, orderData.OrderNumber, nil
}

func saveOrderInfo(txApp core.App, collection *core.Collection, orderData models.OrderData) (string, error) {

	// collection, err := txApp.FindCollectionByNameOrId(constants.TableLiveOrders)
	// if err != nil {
	// 	return "", err
	// }

	record := core.NewRecord(collection)

	record.Set("reference", orderData.Reference)
	record.Set("transaction_type", "L")

	record.Set("order_number", orderData.OrderNumber)

	record.Set("process", "create")
	record.Set("pos_number", "CUSTOMER")

	// hesapla
	record.Set("total_ht", orderData.Total)     //
	record.Set("total_tax", orderData.TotalTax) //
	record.Set("total_ttc", orderData.Total)    //
	record.Set("total_discount", orderData.Reference)
	record.Set("shipping_price", orderData.ShippingPrice)

	//
	record.Set("order_status_id", orderData.OrderStatus)

	record.Set("table_no", orderData.TableNo)

	record.Set("delivery_time", orderData.DeliveryTime)
	// record.Set("delivery_minute", orderData.de)
	record.Set("order_time", orderData.OrderTime)

	//TODO
	record.Set("date", types.NowDateTime().Time().Format("2006-01-02"))
	record.Set("customer", orderData.CustomerID)

	// validate and persist
	// (use SaveNoValidate to skip fields validation)
	// err = txApp.Save(record)
	// if err != nil {
	// 	return "", err
	// }
	if err := txApp.Save(record); err != nil {
		return "", err
	}

	return record.Id, nil
}

func saveOrderProduct(txApp core.App, collection *core.Collection, orderProduct models.OrderProduct, orderID string) (string, error) {

	record := core.NewRecord(collection)

	record.Set("live_order", orderID)
	record.Set("product_id", orderProduct.ProductID)
	record.Set("status", true)

	record.Set("name", orderProduct.Name)
	record.Set("short_name", orderProduct.Name)

	// fmt.Println("save product name", *orderProduct.Name)

	record.Set("unit", "U")
	record.Set("quantity", orderProduct.Quantity)

	record.Set("price_ht", orderProduct.Price)
	record.Set("price_ttc", orderProduct.Price)
	record.Set("total_ht", orderProduct.Total)
	record.Set("total_ttc", orderProduct.Total)
	record.Set("total_tax", orderProduct.Tax)
	record.Set("tax_rate", orderProduct.TaxRate)

	record.Set("prix_revt", nil)
	record.Set("image", nil)

	record.Set("lbl_printer_id", orderProduct.LblPrinterID)
	record.Set("printer_id", orderProduct.PrinterID)
	record.Set("c_sort_order", orderProduct.CSortOrder)
	record.Set("category_id", orderProduct.CategoryID)

	record.Set("created_by", "CUSTOMER")
	record.Set("updated_by", nil)

	if err := txApp.Save(record); err != nil {
		return "", err
	}

	return record.Id, nil

}

func saveOrderProductOptionValue(txApp core.App, collection *core.Collection, optionValue models.OrderOptionValue, orderID string, orderProductID string) error {

	// collection, err := app.FindCollectionByNameOrId(constants.TableLiveOrderProductOptions)
	// if err != nil {
	// 	return err
	// }

	record := core.NewRecord(collection)

	record.Set("live_order", orderID)
	record.Set("live_order_product", orderProductID)

	record.Set("status", true)

	record.Set("option_id", optionValue.OptionID)
	record.Set("option_value_id", optionValue.OptionValueID)

	record.Set("name", optionValue.Name)
	record.Set("short_name", optionValue.Name)

	record.Set("unit", "U")
	record.Set("quantity", optionValue.Quantity)

	record.Set("price_ht", optionValue.Price)
	record.Set("price_ttc", optionValue.Price)

	//hesapla
	record.Set("total_ht", optionValue.TotalHT)
	record.Set("total_ttc", optionValue.TotalTTC)

	record.Set("total_tax", optionValue.Tax)
	record.Set("tax_rate", optionValue.TaxRate)

	record.Set("prix_revt", nil)
	record.Set("image", nil)

	record.Set("title", optionValue.Title)
	record.Set("po_sort_order", optionValue.PoSortOrder)

	record.Set("created_by", "CUSTOMER")
	record.Set("updated_by", nil)

	if err := txApp.Save(record); err != nil {
		return err
	}

	return nil

}
