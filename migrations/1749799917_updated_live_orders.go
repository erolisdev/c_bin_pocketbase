package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1988651086")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation2168032777")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1988651086")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(26, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1751747783",
			"hidden": false,
			"id": "relation2168032777",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "customer",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
