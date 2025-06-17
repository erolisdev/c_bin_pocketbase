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

		log.Println("Creating collection: payments")
		paymentsCollection := core.NewBaseCollection(constants.TablePayments)

		paymentsCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		paymentsCollection.ListRule = types.Pointer("@request.auth.id != ''")
		paymentsCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		paymentsCollection.UpdateRule = types.Pointer("1=2")
		paymentsCollection.DeleteRule = types.Pointer("1=2")

		// paymentsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")
		// paymentsCollection.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.isManager = true")

		// ticketsCollection, err := app.FindCollectionByNameOrId("tickets")
		// if err != nil {
		// 	return err
		// }

		// paymentsCollection.Fields.Add(&core.RelationField{Name: "ticket", CollectionId: ticketsCollection.Id, MaxSelect: 1})

		// returnsCollection, err := app.FindCollectionByNameOrId("returns")
		// if err != nil {
		// 	return err
		// }

		// paymentsCollection.Fields.Add(&core.RelationField{Name: "return_record", CollectionId: returnsCollection.Id, MaxSelect: 1})

		paymentMethodsCollection, err := app.FindCollectionByNameOrId(constants.TablePaymentMethods)
		if err != nil {
			return err
		}

		paymentsCollection.Fields.Add(&core.RelationField{Name: "payment_method", CollectionId: paymentMethodsCollection.Id, MaxSelect: 1})
		paymentsCollection.Fields.Add(&core.SelectField{Name: "transaction_type", Values: []string{"L", "T"}, MaxSelect: 1, Required: true})
		// paymentsCollection.Fields.Add(&core.NumberField{Name: "ticket_no", Required: true})
		paymentsCollection.Fields.Add(&core.TextField{Name: "payment_datetime"})
		paymentsCollection.Fields.Add(&core.TextField{Name: "currency"})
		paymentsCollection.Fields.Add(&core.TextField{Name: "transaction_reference"}) // Card odemelerinin referance numaarasi
		paymentsCollection.Fields.Add(&core.NumberField{Name: "amount", Required: true})
		paymentsCollection.Fields.Add(&core.NumberField{Name: "cash_received"})     //alinan
		paymentsCollection.Fields.Add(&core.NumberField{Name: "cash_change_given"}) // para ustu

		// Go hook Negative
		paymentsCollection.Fields.Add(&core.NumberField{Name: "total_tax_amount_returned", Max: types.Pointer(0.0)}) // Go hook Negative
		paymentsCollection.Fields.Add(&core.NumberField{Name: "total_amount_ttc_returned", Max: types.Pointer(0.0)}) // Go hook Negative
		paymentsCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true})                 // Go hook
		paymentsCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})                  // Go hook

		paymentsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		paymentsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		//indexes
		// paymentsCollection.AddIndex("idx_payments_payment_ticket", false, "ticket", "")
		paymentsCollection.AddIndex("idx_payments_payment_method", false, "payment_method", "")
		// paymentsCollection.AddIndex("idx_payments_return_record", false, "return_record", "")
		// paymentsCollection.AddIndex("idx_payments_ticket_no", false, "ticket_no", "")

		if err := app.Save(paymentsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: payments - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
