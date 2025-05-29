package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: audit_logs")
		auditLogsCollection := core.NewBaseCollection("audit_logs")

		auditLogsCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		auditLogsCollection.ListRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		auditLogsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users'")
		auditLogsCollection.UpdateRule = nil
		auditLogsCollection.DeleteRule = nil

		auditLogsCollection.Fields.Add(&core.TextField{Name: "timestamp", Required: true})

		usersCollection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		auditLogsCollection.Fields.Add(&core.RelationField{Name: "user", CollectionId: usersCollection.Id, MaxSelect: 1})
		auditLogsCollection.Fields.Add(&core.TextField{Name: "action", Required: true})
		auditLogsCollection.Fields.Add(&core.TextField{Name: "details", Required: true})
		auditLogsCollection.Fields.Add(&core.TextField{Name: "previous_record_hash", Required: true}) // Go hook
		auditLogsCollection.Fields.Add(&core.TextField{Name: "current_record_hash", Required: true})  // Go hook

		auditLogsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		auditLogsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		//indexes
		auditLogsCollection.AddIndex("idx_au_logs_action", false, "action", "")
		auditLogsCollection.AddIndex("idx_au_logs_user", false, "user", "")

		if err := app.Save(auditLogsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: audit_logs - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
