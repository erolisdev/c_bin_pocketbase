package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1317801859")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_ZAjuDkHDEP` + "`" + ` ON ` + "`" + `store_zones` + "`" + ` (` + "`" + `city` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_pZdmt7dQ9e` + "`" + ` ON ` + "`" + `store_zones` + "`" + ` (` + "`" + `zone` + "`" + `)"
			],
			"name": "store_zones"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1317801859")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_ZAjuDkHDEP` + "`" + ` ON ` + "`" + `zones` + "`" + ` (` + "`" + `city` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_pZdmt7dQ9e` + "`" + ` ON ` + "`" + `zones` + "`" + ` (` + "`" + `zone` + "`" + `)"
			],
			"name": "zones"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
