package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_537519633")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "",
			"updateRule": ""
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_537519633")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true",
			"updateRule": "@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
