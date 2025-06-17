package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1988651086")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX idx_live_orders_order_number_date ON live_orders (order_number, date)",
				"CREATE INDEX ` + "`" + `idx_order_status_id` + "`" + ` ON ` + "`" + `live_orders` + "`" + ` (order_status_id)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("number504055712")

		// remove field
		collection.Fields.RemoveById("number1577914996")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1988651086")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX idx_live_orders_order_number_date ON live_orders (order_number, date)",
				"CREATE UNIQUE INDEX idx_remote_order_id_date ON live_orders (remote_order_id,date)",
				"CREATE INDEX ` + "`" + `idx_order_status_id` + "`" + ` ON ` + "`" + `live_orders` + "`" + ` (order_status_id)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "number504055712",
			"max": null,
			"min": null,
			"name": "local_order_id",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "number1577914996",
			"max": null,
			"min": null,
			"name": "remote_order_id",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
