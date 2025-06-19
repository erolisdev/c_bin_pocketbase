package eronorhooks

import (
	"c_bin_pocketbase/constants"
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func AuditHook(app *pocketbase.PocketBase) {

	app.OnRecordAfterCreateSuccess(constants.TableLiveOrders).BindFunc(func(e *core.RecordEvent) error {
		// e.App
		// e.Record

		auditCollection, err := e.App.FindCollectionByNameOrId(constants.TableAuditLogs)
		if err != nil {
			fmt.Println("Error finding audit collection: ", err)
			return err
		}

		err = e.App.RunInTransaction(func(txApp core.App) error {
			auditRecord := core.NewRecord(auditCollection)

			timeStamp := time.Now().UnixMilli()

			auditRecord.Set("timestamp", timeStamp)
			auditRecord.Set("user", e.Record.GetString("created_by"))
			auditRecord.Set("action", "create")
			auditRecord.Set("details", "Test Data")
			auditRecord.Set("previous_record_hash", "hash1")
			auditRecord.Set("current_record_hash", "hash2")

			if err := txApp.Save(auditRecord); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			fmt.Println("Error creating audit record: ", err)
			return err
		}

		return e.Next()
	})

}
