package migrations

// import (
// 	"log"

// 	"github.com/pocketbase/pocketbase/core"
// 	m "github.com/pocketbase/pocketbase/migrations"
// 	"github.com/pocketbase/pocketbase/tools/types"
// )

// func init() {
// 	m.Register(func(app core.App) error {

// 		log.Println("Creating collection: reporting_ticket_line_options")
// 		reportingTicketLineOptions := core.NewBaseCollection("reporting_ticket_line_options")

// 		reportingTicketLineOptions.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
// 		reportingTicketLineOptions.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
// 		reportingTicketLineOptions.UpdateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
// 		// reportingTicketLineOptions.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.isManager = true")

// 		// ticket relation
// 		ticketsCollection, err := app.FindCollectionByNameOrId("tickets")
// 		if err != nil {
// 			return err
// 		}
// 		reportingTicketLineOptions.Fields.Add(&core.RelationField{Name: "ticket", CollectionId: ticketsCollection.Id, MaxSelect: 1})

// 		// ticket item relation
// 		ticketItemsCollection, err := app.FindCollectionByNameOrId("ticket_lines")
// 		if err != nil {
// 			return err
// 		}
// 		reportingTicketLineOptions.Fields.Add(&core.RelationField{Name: "ticket_line", CollectionId: ticketItemsCollection.Id, MaxSelect: 1})

// 		// option relation
// 		optionValuesCollection, err := app.FindCollectionByNameOrId("option_values")
// 		if err != nil {
// 			return err
// 		}
// 		reportingTicketLineOptions.Fields.Add(&core.RelationField{Name: "option_value", CollectionId: optionValuesCollection.Id, MaxSelect: 1})

// 		//
// 		reportingTicketLineOptions.Fields.Add(&core.TextField{Name: "name"})
// 		reportingTicketLineOptions.Fields.Add(&core.NumberField{Name: "price_ttc"})
// 		reportingTicketLineOptions.Fields.Add(&core.TextField{Name: "unit"})
// 		reportingTicketLineOptions.Fields.Add(&core.NumberField{Name: "quantity_unit"})
// 		reportingTicketLineOptions.Fields.Add(&core.NumberField{Name: "total_ttc"})

// 		reportingTicketLineOptions.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
// 		reportingTicketLineOptions.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

// 		// indexs
// 		reportingTicketLineOptions.AddIndex("idx_re_op_ticket", false, "ticket", "")
// 		reportingTicketLineOptions.AddIndex("idx_re_op_ticket_line", false, "ticket_line", "")
// 		reportingTicketLineOptions.AddIndex("idx_re_op_option_value", false, "option_value", "")

// 		if err := app.Save(reportingTicketLineOptions); err != nil {
// 			return err
// 		}

// 		log.Println("--- Migration: reporting ticket item options - Tamamlandı ---")
// 		return nil
// 	}, func(app core.App) error {
// 		// Revert
// 		return nil
// 	})
// }
