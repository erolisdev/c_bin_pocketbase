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

		log.Println("Creating collection: live_order_products")
		liveOrderProductsCollection := core.NewBaseCollection(constants.TableLiveOrderProducts)

		liveOrderProductsCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		liveOrderProductsCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		liveOrderProductsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")

		liveOrdersCollection, err := app.FindCollectionByNameOrId(constants.TableLiveOrders)
		if err != nil {
			return err
		}

		liveOrderProductsCollection.Fields.Add(&core.RelationField{Name: "live_order", CollectionId: liveOrdersCollection.Id, MaxSelect: 1})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "local_order_id"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "product_id", Required: true})
		liveOrderProductsCollection.Fields.Add(&core.BoolField{Name: "status"}) // iptal edilenler
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "name", Required: true})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "short_name"})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "unit"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "quantity"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "price_ht"})  // price of product
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "price_ttc"}) // price of product
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "total_ht"})  // price with options
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "total_ttc"}) // price with options
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "total_tax"}) // total_tax

		taxesCollection, err := app.FindCollectionByNameOrId(constants.TableTaxRates)
		if err != nil {
			return err
		}

		liveOrderProductsCollection.Fields.Add(&core.RelationField{Name: "tax_rate", CollectionId: taxesCollection.Id, MaxSelect: 1, MinSelect: 1}) // tax rate //TODO
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "prix_revt"})

		liveOrdersCollection.Fields.Add(&core.TextField{Name: "product_discount"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "product_discount_rate"})

		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "image"})

		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_id", Required: true})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "template_id", Required: true})

		liveOrdersCollection.Fields.Add(&core.TextField{Name: "label_printer_id"})

		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "c_sort_order"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "category_id"})

		//

		// customersCollection, err := app.FindCollectionByNameOrId(constants.TableCustomers)
		// if err != nil {
		// 	return err
		// }

		// liveOrderProductsCollection.Fields.Add(&core.RelationField{Name: "customer", CollectionId: customersCollection.Id, MaxSelect: 1})

		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "created_by"})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "updated_by"})

		liveOrderProductsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		liveOrderProductsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		// indexs
		liveOrderProductsCollection.AddIndex("idx_live_order_products_idx_order", false, "live_order", "")

		if err := app.Save(liveOrderProductsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: live_order_products- Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
