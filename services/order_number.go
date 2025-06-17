package services

import (
	"c_bin_pocketbase/constants"
	"fmt"
	"strings"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func GetOrderNumber(app *pocketbase.PocketBase) (*int, error) {
	orderNumber := 1

	var lastSavedRecord = core.Record{}
	const customLayout = "2006-01-02 15:04:05.000Z07:00"

	now := time.Now().UTC()
	const startHour = 5

	// İlgili günlerin sabah 5'ini önceden hesaplayalım.
	todayAt5 := time.Date(now.Year(), now.Month(), now.Day(), startHour, 0, 0, 0, time.UTC)

	var startTime, endTime time.Time

	// Eğer mevcut zaman, bugünün başlangıç saatinden (sabah 5) önceyse...
	if now.Before(todayAt5) {
		// Periyot dün sabah 5'te başladı.
		startTime = endTime.AddDate(0, 0, -1) // Bitiş saatinden 1 gün çıkar.
		endTime = todayAt5
	} else {
		// Periyot bu sabah 5'te başladı.
		startTime = todayAt5
		endTime = startTime.AddDate(0, 0, 1) // Başlangıç saatine 1 gün ekle.
	}

	fmt.Println("startTime", startTime, "start format", startTime.Format(customLayout))

	err := app.RecordQuery(constants.TableLiveOrders).
		Where(dbx.Between("created", startTime.Format(customLayout), endTime.Format(customLayout))).
		OrderBy("order_number DESC").
		One(&lastSavedRecord)

	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, fmt.Errorf("live_orders save order number: '%s'", err.Error())
	}

	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		orderNumber = 1
	} else {
		orderNumber = lastSavedRecord.GetInt("order_number") + 1

	}

	return &orderNumber, nil
}
