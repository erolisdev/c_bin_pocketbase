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

		log.Println("Creating collection: returns")
		returnsCollection := core.NewBaseCollection(constants.TableReturns)

		returnsCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		returnsCollection.ListRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		returnsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")

		ticketsCollection, err := app.FindCollectionByNameOrId(constants.TableTickets)
		if err != nil {
			return err
		}

		returnsCollection.Fields.Add(&core.RelationField{Name: "original_ticket", CollectionId: ticketsCollection.Id, MaxSelect: 1})
		returnsCollection.Fields.Add(&core.NumberField{Name: "ticket_no", Required: true})
		
		returnsCollection.Fields.Add(&core.SelectField{Name: "transaction_type", Values: []string{"L", "T"}, MaxSelect: 1, Required: true})
		returnsCollection.Fields.Add(&core.SelectField{Name: "return_type", Required: true, Values: []string{"RETURN", "CANCELLATION", "OTHER"}, MaxSelect: 1})
		returnsCollection.Fields.Add(&core.TextField{Name: "return_datetime"})

		returnsCollection.Fields.Add(&core.NumberField{Name: "total_ht", Max: types.Pointer(0.0), Required: true})  // Go hook Negative
		returnsCollection.Fields.Add(&core.NumberField{Name: "total_ttc", Max: types.Pointer(0.0), Required: true}) // Go hook Negative
		returnsCollection.Fields.Add(&core.NumberField{Name: "total_tax", Max: types.Pointer(0.0), Required: true}) // Go hook Negative

		returnsCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true}) // Go hook
		returnsCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})  // Go hook

		// customersCollection, err := app.FindCollectionByNameOrId(constants.TableCustomers)
		// if err != nil {
		// 	return err
		// }

		// returnsCollection.Fields.Add(&core.RelationField{Name: "customer", CollectionId: customersCollection.Id, MaxSelect: 1})

		returnsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		returnsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		returnsCollection.AddIndex("idx_returns_original_ticket", true, "original_ticket", "")
		returnsCollection.AddIndex("idx_returns_ticket_no", false, "ticket_no", "")

		if err := app.Save(returnsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: returns - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
