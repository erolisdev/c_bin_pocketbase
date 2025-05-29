package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: tickets")
		ticketsCollection := core.NewBaseCollection("tickets")

		// ticketsCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		// ticketsCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		// ticketsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")

		ticketsCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		ticketsCollection.ListRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		ticketsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")

		ticketsCollection.Fields.Add(&core.NumberField{Name: "ticket_id", Required: true, Min: types.Pointer(1.0), Presentable: true})
		ticketsCollection.Fields.Add(&core.NumberField{Name: "ticket_no", Required: true, Min: types.Pointer(1.0)})
		ticketsCollection.Fields.Add(&core.NumberField{Name: "order_number", Required: true, Min: types.Pointer(1.0)})
		ticketsCollection.Fields.Add(&core.NumberField{Name: "web_order_number"})
		ticketsCollection.Fields.Add(&core.TextField{Name: "ticket_datetime"})
		// ticketsCollection.Fields.Add(&core.TextField{Name: "created_by"})
		ticketsCollection.Fields.Add(&core.TextField{Name: "caissier"})                                                    // name of the cashier
		ticketsCollection.Fields.Add(&core.NumberField{Name: "total_amount_ht", Min: types.Pointer(0.0), Required: true})  // Go hook
		ticketsCollection.Fields.Add(&core.NumberField{Name: "total_tax_amount", Min: types.Pointer(0.0), Required: true}) // Go hook
		ticketsCollection.Fields.Add(&core.NumberField{Name: "total_amount_ttc", Min: types.Pointer(0.0), Required: true}) // Go hook
		ticketsCollection.Fields.Add(&core.TextField{Name: "previous_record_hash"})                                        // Go hook
		ticketsCollection.Fields.Add(&core.TextField{Name: "current_record_hash"})                                         // Go hook
		ticketsCollection.Fields.Add(&core.TextField{Name: "model"})                                                       // barcode

		customersCollection, err := app.FindCollectionByNameOrId("customers")
		if err != nil {
			return err
		}

		ticketsCollection.Fields.Add(&core.RelationField{Name: "customer", CollectionId: customersCollection.Id, MaxSelect: 1})
		ticketsCollection.Fields.Add(&core.TextField{Name: "customer_fullname"})

		ticketsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		ticketsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		ticketsCollection.AddIndex("idx_ticket_id", true, "ticket_id", "")
		ticketsCollection.AddIndex("idx_ticket_no", false, "ticket_no", "")

		if err := app.Save(ticketsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: tickets - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
