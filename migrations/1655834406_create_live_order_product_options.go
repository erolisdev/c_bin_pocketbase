package migrations

import (
	"c_bin_pocketbase/constants"
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: live_order_product_options")
		liveOrderProductOptionsCollection := core.NewBaseCollection(constants.TableLiveOrderProductOptions)

		liveOrderProductOptionsCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		liveOrderProductOptionsCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		liveOrderProductOptionsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")
		// liveOrderProductOptionsCollection.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.isManager = true")

		liveOrdersCollection, err := app.FindCollectionByNameOrId(constants.TableLiveOrders)
		if err != nil {
			return err
		}

		liveOrderProductOptionsCollection.Fields.Add(&core.RelationField{Name: "live_order", CollectionId: liveOrdersCollection.Id, MaxSelect: 1})

		liveOrderProductsCollection, err := app.FindCollectionByNameOrId(constants.TableLiveOrderProducts)
		if err != nil {
			return err
		}

		liveOrderProductOptionsCollection.Fields.Add(&core.RelationField{Name: "live_order_product", CollectionId: liveOrderProductsCollection.Id, MaxSelect: 1})

		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "local_order_id"})
		liveOrderProductOptionsCollection.Fields.Add(&core.BoolField{Name: "status"}) // iptal edilenler
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "product_id"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "option_id"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "option_value_id"})
		liveOrderProductOptionsCollection.Fields.Add(&core.TextField{Name: "name", Required: true})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "short_name"})

		liveOrderProductOptionsCollection.Fields.Add(&core.TextField{Name: "unit"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "quantity"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "price_ht"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "price_ttc"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "total_ht"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "total_ttc"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "total_tax"})

		taxesCollection, err := app.FindCollectionByNameOrId(constants.TableTaxRates)
		if err != nil {
			return err
		}

		liveOrderProductOptionsCollection.Fields.Add(&core.RelationField{Name: "tax_rate", CollectionId: taxesCollection.Id, MaxSelect: 1, MinSelect: 1}) // tax rate //TODO
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "prix_revt"})

		liveOrderProductOptionsCollection.Fields.Add(&core.TextField{Name: "created_by"})
		liveOrderProductOptionsCollection.Fields.Add(&core.TextField{Name: "updated_by"})

		liveOrderProductOptionsCollection.Fields.Add(&core.TextField{Name: "image"})

		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_id", Required: true})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "template_id", Required: true})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "label_printer_id"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "title"})
		liveOrderProductOptionsCollection.Fields.Add(&core.NumberField{Name: "po_sort_order"})

		// liveOrdersCollection.Fields.Add(&core.TextField{Name: "printer_ip", Required: true})
		// liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_port"})

		// liveOrdersCollection.Fields.Add(&core.TextField{Name: "product_discount"})
		// liveOrdersCollection.Fields.Add(&core.TextField{Name: "product_discount_rate"})
		//

		// customersCollection, err := app.FindCollectionByNameOrId(constants.TableCustomers)
		// if err != nil {
		// 	return err
		// }

		// liveOrderProductOptionsCollection.Fields.Add(&core.RelationField{Name: "customer", CollectionId: customersCollection.Id, MaxSelect: 1})

		liveOrderProductOptionsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		liveOrderProductOptionsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		// indexs
		liveOrderProductOptionsCollection.AddIndex("idx_l_order_product_options_order", false, "live_order", "")
		liveOrderProductOptionsCollection.AddIndex("idx_l_order_product_options_product", false, "live_order_product", "")

		if err := app.Save(liveOrderProductOptionsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: live_order_product_options- Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
