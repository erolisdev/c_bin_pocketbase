package utils

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func GetRequestUserIDAlias(authRecord *core.Record) string {

	if authRecord == nil {
		return "GUEST"
	}

	if authRecord.Collection().Name == "customers" {
		return fmt.Sprint("C-", authRecord.Id)
	}

	if authRecord.Collection().Name == "users" {
		return fmt.Sprint("U-", authRecord.Id)
	}

	return "SYSTEM"
}
