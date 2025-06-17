package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3611327082")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(19, []byte(`{
			"hidden": false,
			"id": "number2051285377",
			"max": null,
			"min": null,
			"name": "lbl_printer_id",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(20, []byte(`{
			"hidden": false,
			"id": "number1189890378",
			"max": null,
			"min": null,
			"name": "printer_id",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(21, []byte(`{
			"hidden": false,
			"id": "number1769643497",
			"max": null,
			"min": null,
			"name": "c_sort_order",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(22, []byte(`{
			"hidden": false,
			"id": "number306617826",
			"max": null,
			"min": null,
			"name": "category_id",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3611327082")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("number2051285377")

		// remove field
		collection.Fields.RemoveById("number1189890378")

		// remove field
		collection.Fields.RemoveById("number1769643497")

		// remove field
		collection.Fields.RemoveById("number306617826")

		return app.Save(collection)
	})
}
