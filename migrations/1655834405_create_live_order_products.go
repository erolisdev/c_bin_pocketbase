package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: live_order_products")
		liveOrderProductsCollection := core.NewBaseCollection("live_order_products")

		liveOrderProductsCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		liveOrderProductsCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		liveOrderProductsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")

		liveOrdersCollection, err := app.FindCollectionByNameOrId("live_orders")
		if err != nil {
			log.Println("Creating collection: live_order_products 1_1")
			return err
		}

		liveOrderProductsCollection.Fields.Add(&core.RelationField{Name: "live_order", CollectionId: liveOrdersCollection.Id, MaxSelect: 1})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "local_order_id"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "product_id", Required: true})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "unit"})
		liveOrderProductsCollection.Fields.Add(&core.BoolField{Name: "status"}) // iptal edilenler
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "quantity_unit"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "uprice_w_tax"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "uprice_standart"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "prix_revt"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "total_row_tax"})
		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "total_row_w_tax"}) //TODO

		taxesCollection, err := app.FindCollectionByNameOrId("tax_rates")
		if err != nil {
			return err
		}

		log.Println("Creating collection: live_order_products 2")

		liveOrderProductsCollection.Fields.Add(&core.RelationField{Name: "tax_rate", CollectionId: taxesCollection.Id, MaxSelect: 1, MinSelect: 1}) // tax rate //TODO

		liveOrderProductsCollection.Fields.Add(&core.NumberField{Name: "product_id"})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "created_by"})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "updated_by"})

		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "name"})
		liveOrderProductsCollection.Fields.Add(&core.TextField{Name: "image"})

		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_id", Required: true})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "printer_ip", Required: true})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_port"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "template_id", Required: true})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "label_printer_id"})

		liveOrdersCollection.Fields.Add(&core.TextField{Name: "product_discount"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "product_discount_rate"})
		//

		customersCollection, err := app.FindCollectionByNameOrId("customers")
		if err != nil {
			return err
		}

		liveOrderProductsCollection.Fields.Add(&core.RelationField{Name: "customer", CollectionId: customersCollection.Id, MaxSelect: 1})

		liveOrderProductsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		liveOrderProductsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		// indexs
		liveOrderProductsCollection.AddIndex("idx_live_order_products_idx_order", false, "live_order", "")

		log.Println("Creating collection: live_order_products 3")

		if err := app.Save(liveOrderProductsCollection); err != nil {
			log.Println("Creating collection: live_order_products 4")
			return err
		}

		log.Println("--- Migration: live_order_products- Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
