package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2854637623")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_product_id` + "`" + ` ON ` + "`" + `store_products` + "`" + ` (product_id)",
				"CREATE INDEX ` + "`" + `idx_category_id` + "`" + ` ON ` + "`" + `store_products` + "`" + ` (category)"
			]
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_175481600",
			"hidden": false,
			"id": "relation105650625",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "category",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2854637623")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_product_id` + "`" + ` ON ` + "`" + `store_products` + "`" + ` (product_id)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_category_id` + "`" + ` ON ` + "`" + `store_products` + "`" + ` (category)"
			]
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_175481600",
			"hidden": false,
			"id": "relation105650625",
			"maxSelect": 999,
			"minSelect": 0,
			"name": "category",
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
