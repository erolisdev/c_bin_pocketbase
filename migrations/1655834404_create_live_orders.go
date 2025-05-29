package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: live_orders")
		liveOrdersCollection := core.NewBaseCollection("live_orders")

		liveOrdersCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		liveOrdersCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		liveOrdersCollection.UpdateRule = types.Pointer("@request.auth.id != ''")

		// liveOrdersCollection.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.isManager = true")

		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "local_order_id"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "order_number"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "remote_order_id"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "web_order_number"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "process"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "pos_number"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "total_tax", Min: types.Pointer(0.0), Required: true})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "total_w_tax", Min: types.Pointer(0.0), Required: true})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "total_discount"})
		liveOrdersCollection.Fields.Add(&core.NumberField{Name: "remainder"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "caissier"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "order_status_id"})

		liveOrdersCollection.Fields.Add(&core.TextField{Name: "table_no"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "delivery_time"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "delivery_minute"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "order_time"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "order_customer"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "date"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "name"})
		liveOrdersCollection.Fields.Add(&core.TextField{Name: "shipping_firstname"}) //customer name

		// merged orders
		liveOrdersCollection.Fields.Add(&core.BoolField{Name: "is_merged"})
		liveOrdersCollection.Fields.Add(&core.JSONField{Name: "merged_orders"})
		liveOrdersCollection.Fields.Add(&core.JSONField{Name: "merged_local_ids"})

		customersCollection, err := app.FindCollectionByNameOrId("customers")
		if err != nil {
			return err
		}

		liveOrdersCollection.Fields.Add(&core.RelationField{Name: "customer", CollectionId: customersCollection.Id, MaxSelect: 1})

		liveOrdersCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		liveOrdersCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		// liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_id"})
		// liveOrdersCollection.Fields.Add(&core.TextField{Name: "printer_ip", Required: true})
		// liveOrdersCollection.Fields.Add(&core.NumberField{Name: "printer_port"})
		// liveOrdersCollection.Fields.Add(&core.NumberField{Name: "template_id", Required: true})
		// liveOrdersCollection.Fields.Add(&core.TextField{Name: "label_printer_id", Required: true})

		// indexs
		liveOrdersCollection.Indexes = []string{"CREATE UNIQUE INDEX idx_live_orders_order_number_date ON live_orders (order_number, date)"}
		liveOrdersCollection.Indexes = append(liveOrdersCollection.Indexes, "CREATE UNIQUE INDEX idx_remote_order_id_date ON live_orders (remote_order_id,date)")
		liveOrdersCollection.AddIndex("idx_order_status_id", false, "order_status_id", "")

		if err := app.Save(liveOrdersCollection); err != nil {
			return err
		}

		log.Println("--- Migration: live_orders- Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
