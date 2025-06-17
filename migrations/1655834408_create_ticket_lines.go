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

		log.Println("Creating collection: ticket_lines")
		ticketLinesCollection := core.NewBaseCollection(constants.TableTicketLines)

		ticketLinesCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		ticketLinesCollection.ListRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		ticketLinesCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")

		ticketsCollection, err := app.FindCollectionByNameOrId(constants.TableTickets)
		if err != nil {
			return err
		}

		ticketLinesCollection.Fields.Add(&core.RelationField{Name: "ticket", CollectionId: ticketsCollection.Id, MaxSelect: 1})

		productsCollection, err := app.FindCollectionByNameOrId(constants.TableStoreProducts)
		if err != nil {
			return err
		}

		ticketLinesCollection.Fields.Add(&core.RelationField{Name: "product", CollectionId: productsCollection.Id, MaxSelect: 1, Required: true})
		ticketLinesCollection.Fields.Add(&core.SelectField{Name: "transaction_type", Values: []string{"L", "T"}, MaxSelect: 1, Required: true})
		ticketLinesCollection.Fields.Add(&core.TextField{Name: "unit", Required: true})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "quantity", Required: true, Min: types.Pointer(0.0)})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "price_ht", Required: true})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "price_ttc", Required: true})
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "tax_rate"})
		ticketLinesCollection.Fields.Add(&core.JSONField{Name: "chosen_options"})

		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_ht", Min: types.Pointer(0.0), Required: true})  // Go hook
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_ttc", Min: types.Pointer(0.0), Required: true}) // Go hook
		ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_tax", Min: types.Pointer(0.0), Required: true}) // Go hook
		// ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_tax_amount", Min: types.Pointer(0.0), Required: true}) // Go hook
		// ticketLinesCollection.Fields.Add(&core.NumberField{Name: "total_amount_ttc", Min: types.Pointer(0.0), Required: true}) // Go hook
		ticketLinesCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true}) // Go hook
		ticketLinesCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})  // Go hook

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
