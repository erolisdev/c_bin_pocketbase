package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_4242123432")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(17, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text4145232419",
			"max": 0,
			"min": 0,
			"name": "day_reference",
			"pattern": "^([01]\\d|2[0-3]):([0-5]\\d)$",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_4242123432")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(17, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text4145232419",
			"max": 0,
			"min": 0,
			"name": "day_reference",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
