package eronorhooks

// func OrderNumberCreate(app *pocketbase.PocketBase) {

// 	// fires only for "users" and "articles" records
// 	app.OnRecordCreateExecute(constants.TableLiveOrders).BindFunc(func(e *core.RecordEvent) error {
// e.App
// e.Record

// var lastSavedRecord = core.Record{}
// const customLayout = "2006-01-02 15:04:05.000Z07:00"

// now := time.Now().UTC()
// const startHour = 5

// // İlgili günlerin sabah 5'ini önceden hesaplayalım.
// todayAt5 := time.Date(now.Year(), now.Month(), now.Day(), startHour, 0, 0, 0, time.UTC)

// var startTime, endTime time.Time

// // Eğer mevcut zaman, bugünün başlangıç saatinden (sabah 5) önceyse...
// if now.Before(todayAt5) {
// 	// Periyot dün sabah 5'te başladı.
// 	startTime = endTime.AddDate(0, 0, -1) // Bitiş saatinden 1 gün çıkar.
// 	endTime = todayAt5
// } else {
// 	// Periyot bu sabah 5'te başladı.
// 	startTime = todayAt5
// 	endTime = startTime.AddDate(0, 0, 1) // Başlangıç saatine 1 gün ekle.
// }

// fmt.Println("startTime", startTime, "start format", startTime.Format(customLayout))

// err := app.RecordQuery(constants.TableLiveOrders).
// 	Where(dbx.Between("created", startTime.Format(customLayout), endTime.Format(customLayout))).
// 	OrderBy("order_number DESC").
// 	One(&lastSavedRecord)

// if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
// 	return fmt.Errorf("live_orders save order number: '%s'", err.Error())
// }

// if err != nil && strings.Contains(err.Error(), "no rows in result set") {
// 	e.Record.Set("order_number", 1)
// } else {
// 	e.Record.Set("order_number", lastSavedRecord.GetInt("order_number")+1)
// }

// 		return e.Next()
// 	})

// }
