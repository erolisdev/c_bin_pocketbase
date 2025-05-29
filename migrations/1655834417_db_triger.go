package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {

		_, errDelete := app.DB().NewQuery("CREATE TRIGGER IF NOT EXISTS prevent_delete BEFORE DELETE ON audit_logs BEGIN SELECT RAISE(ABORT, 'Bu kayitlar silinemez!'); END;").Execute()

		if errDelete != nil {
			log.Println("Delete trigger not created:", errDelete)
			return errDelete
		}

		_, errUpdate := app.DB().NewQuery("CREATE TRIGGER IF NOT EXISTS prevent_update BEFORE UPDATE ON audit_logs BEGIN SELECT RAISE(ABORT, 'Bu kayıtlar degistirilemez! '); END;").Execute()

		if errUpdate != nil {
			log.Println("Update trigger not created:", errUpdate)
			return errUpdate
		}

		return nil

	}, nil)
}
