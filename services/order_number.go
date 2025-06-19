package services

import (
	"c_bin_pocketbase/constants"
	"c_bin_pocketbase/models"
	"fmt"
	"strings"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func GetOrderNumber(txApp core.App) (*int, error) {
	orderNumber := 1

	var lastSavedRecord = core.Record{}
	const customLayout = "2006-01-02 15:04:05.000Z07:00"
	var startHour, startMinute int
	var startTime, endTime time.Time
	now := time.Now().UTC()

	dayRef := txApp.Store().Get("day_referece_time")

	if dayRef != nil {
		parsedTime, err := time.Parse("15:04", dayRef.(string))
		if err == nil {
			startHour = parsedTime.Hour()
			startMinute = parsedTime.Minute()
		} else {
			startHour = 5
			startMinute = 0
		}
	} else {
		// DB'den değer dönmezse fallback değer ata
		startHour = 5
		startMinute = 0
	}

	// İlgili günlerin sabah 5'ini önceden hesaplayalım.
	todayAtDayRef := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMinute, 0, 0, time.UTC)

	// Eğer mevcut zaman, bugünün başlangıç saatinden (sabah 5) önceyse...
	if now.Before(todayAtDayRef) {
		// Periyot dün sabah 5'te başladı.
		startTime = endTime.AddDate(0, 0, -1) // Bitiş saatinden 1 gün çıkar.
		endTime = todayAtDayRef
	} else {
		// Periyot bu sabah 5'te başladı.
		startTime = todayAtDayRef
		endTime = startTime.AddDate(0, 0, 1) // Başlangıç saatine 1 gün ekle.
	}

	err := txApp.RecordQuery(constants.TableLiveOrders).
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

// onServe de yani app ilk acildiginda set edilecek
func SaveDayReference(app *pocketbase.PocketBase) {
	setting := models.StoreSetting{}

	err := app.DB().
		NewQuery("SELECT day_reference FROM store_settings").
		One(&setting)

	if err == nil {
		if setting.DayRef != "" {
			fmt.Println("set day ref: ", setting.DayRef)
			app.Store().Set("day_referece_time", setting.DayRef)
		}
	}

}
