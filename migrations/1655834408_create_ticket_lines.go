package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: ticket_lines")
		ticketLinesCollection := core.NewBaseCollection("ticket_lines")

		ticketLinesCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		ticketLinesCollection.ListRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		ticketLinesCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")

		ticketsCollection, err := app.FindCollectionByNameOrId("tickets")
		if err != nil {
			return err
		}

		ticketLinesCollection.Fields.Add(&core.RelationField{Name: "ticket", CollectionId: ticketsCollection.Id, MaxSelect: 1})

		productsCollection, err := app.FindCollectionByNameOrId("products")
		if err != nil {
			return err
		}

		ticketLinesCollection.Fields.Add(&core.RelationField{Name: "product", CollectionId: productsCollection.Id, MaxSelect: 1})

		ticketLinesCollection.Fields.Add(&core.TextField{Name: "unit", Required: true})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "quantity_unit", Required: true, Min: types.Pointer(0.0)})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "base_price_ht", Required: true})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "base_tax_rate"})
		ticketLinesCollection.Fields.Add(&core.JSONField{Name: "chosen_options"})

		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "calculated_total_ht", Min: types.Pointer(0.0), Required: true})  // Go hook
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "calculated_total_tax", Min: types.Pointer(0.0), Required: true}) // Go hook
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "calculated_total_ttc", Min: types.Pointer(0.0), Required: true}) // Go hook
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_tax_amount", Min: types.Pointer(0.0), Required: true})     // Go hook
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_amount_ttc", Min: types.Pointer(0.0), Required: true})     // Go hook
		ticketLinesCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true})                            // Go hook
		ticketLinesCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})                             // Go hook

		ticketLinesCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		ticketLinesCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		ticketLinesCollection.AddIndex("idx_ticket", false, "ticket", "")
		ticketLinesCollection.AddIndex("idx_product", false, "product", "")

		if err := app.Save(ticketLinesCollection); err != nil {
			return err
		}

		log.Println("--- Migration: ticket_lines - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
