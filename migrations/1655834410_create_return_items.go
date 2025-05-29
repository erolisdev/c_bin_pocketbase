package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: return_items")
		// parca iadeler icin urun iadeleri icin
		returnItemsCollection := core.NewBaseCollection("return_items")

		returnItemsCollection.ViewRule = types.Pointer("@request.auth.id != ''")
		returnItemsCollection.CreateRule = types.Pointer("@request.auth.id != ''")
		returnItemsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")
		// returnItemsCollection.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.isManager = true")

		returnsCollection, err := app.FindCollectionByNameOrId("returns")
		if err != nil {
			return err
		}

		returnItemsCollection.Fields.Add(&core.RelationField{Name: "return_record", CollectionId: returnsCollection.Id, MaxSelect: 1})

		saleItemsCollection, err := app.FindCollectionByNameOrId("ticket_lines")
		if err != nil {
			return err
		}

		returnItemsCollection.Fields.Add(&core.RelationField{Name: "original_ticket_line", CollectionId: saleItemsCollection.Id, MaxSelect: 1})

		returnItemsCollection.Fields.Add(&core.TextField{Name: "unit", Required: true})
		returnItemsCollection.Fields.Add(&core.NumberField{Name: "quantity_unit_returned", Required: true})
		returnItemsCollection.Fields.Add(&core.NumberField{Name: "total_ht_returned", Required: true})  // Go hook
		returnItemsCollection.Fields.Add(&core.NumberField{Name: "total_tax_returned", Required: true}) // Go hook
		returnItemsCollection.Fields.Add(&core.NumberField{Name: "total_ttc_returned", Required: true}) // Go hook
		returnItemsCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true}) // Go hook
		returnItemsCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})  // Go hook

		returnItemsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		returnItemsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		returnItemsCollection.AddIndex("idx_ritems_return_record", false, "return_record", "")
		returnItemsCollection.AddIndex("idx_ritems_ticket_line", false, "original_ticket_line", "")

		if err := app.Save(returnItemsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: return_items - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
