package eronorhooks

import (
	"github.com/pocketbase/pocketbase"
)

func EronorHooks(app *pocketbase.PocketBase) {
	// ProtectionForDeleteHooks(app)
	AuditHook(app)

}
