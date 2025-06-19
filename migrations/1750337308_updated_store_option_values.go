package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_417461541")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(13, []byte(`{
			"hidden": false,
			"id": "number1352515405",
			"max": null,
			"min": null,
			"name": "reset",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(15, []byte(`{
			"hidden": false,
			"id": "number2063623452",
			"max": null,
			"min": null,
			"name": "status",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(16, []byte(`{
			"hidden": false,
			"id": "number1476078466",
			"max": null,
			"min": null,
			"name": "not_recommanded",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(17, []byte(`{
			"hidden": false,
			"id": "number2128919991",
			"max": null,
			"min": null,
			"name": "image_type",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_417461541")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(13, []byte(`{
			"hidden": false,
			"id": "number1352515405",
			"max": null,
			"min": null,
			"name": "reset",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(15, []byte(`{
			"hidden": false,
			"id": "number2063623452",
			"max": null,
			"min": null,
			"name": "status",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(16, []byte(`{
			"hidden": false,
			"id": "number1476078466",
			"max": null,
			"min": null,
			"name": "not_recommanded",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(17, []byte(`{
			"hidden": false,
			"id": "number2128919991",
			"max": null,
			"min": null,
			"name": "image_type",
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
