package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: reports")
		closuresReportsCollection := core.NewBaseCollection("closures_reports")

		closuresReportsCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		closuresReportsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		closuresReportsCollection.UpdateRule = nil
		closuresReportsCollection.DeleteRule = nil
		// closuresReportsCollection.UpdateRule = types.Pointer("@request.auth.id != ''")
		// closuresReportsCollection.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.isManager = true")

		closuresReportsCollection.Fields.Add(&core.TextField{Name: "timestamp", Required: true})

		usersCollection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		closuresReportsCollection.Fields.Add(&core.RelationField{Name: "user", CollectionId: usersCollection.Id, MaxSelect: 1})
		closuresReportsCollection.Fields.Add(&core.TextField{Name: "action", Required: true})
		closuresReportsCollection.Fields.Add(&core.TextField{Name: "details", Required: true})
		closuresReportsCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true}) // Go hook
		closuresReportsCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})  // Go hook

		closuresReportsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		closuresReportsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		//indexes
		// auditLogsCollection.AddIndex("idx_au_logs_action", false, "action", "")
		// auditLogsCollection.AddIndex("idx_au_logs_user", false, "user", "")

		if err := app.Save(closuresReportsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: closures_reports_collection - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
