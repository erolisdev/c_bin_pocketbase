package eronorhooks

import (
	"c_bin_pocketbase/constants"
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func ProtectionForDeleteHooks(app *pocketbase.PocketBase) {
	app.OnCollectionDelete(constants.TableAuditLogs, constants.TableClosureReports, constants.TableTickets, constants.TableTicketLines, constants.TablePayments).BindFunc(func(e *core.CollectionEvent) error {
		return fmt.Errorf("Delete not allowed %s", e.Collection.Name)
	})

	app.OnRecordDelete(constants.TableAuditLogs, constants.TableClosureReports, constants.TableTickets, constants.TableTicketLines, constants.TablePayments).BindFunc(func(e *core.RecordEvent) error {
		return fmt.Errorf("Delete not allowed %s", e.Record.Collection().Name)
	})

}
