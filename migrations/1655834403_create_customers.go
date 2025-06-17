package migrations

import (
	"c_bin_pocketbase/constants"
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		log.Println("Creating collection: auth_customers")
		customersCollection := core.NewAuthCollection(constants.TableCustomers)

		customersCollection.ViewRule = types.Pointer("@request.auth.id != '' && @request.auth.id = id")
		customersCollection.UpdateRule = types.Pointer("@request.auth.id != '' && @request.auth.id = id")
		customersCollection.DeleteRule = types.Pointer("@request.auth.id != '' && @request.auth.id = id")
		customersCollection.CreateRule = types.Pointer("")
		customersCollection.AuthRule = types.Pointer("verified = true")

		customersCollection.Fields.Add(&core.TextField{Name: "firstname", Required: true})
		customersCollection.Fields.Add(&core.TextField{Name: "lastname", Required: true})
		customersCollection.Fields.Add(&core.EmailField{Name: "email", Required: true})
		customersCollection.Fields.Add(&core.TextField{Name: "telephone", Required: true})
		customersCollection.Fields.Add(&core.TextField{Name: "company"})
		customersCollection.Fields.Add(&core.JSONField{Name: "address"}) //json address list
		customersCollection.Fields.Add(&core.TextField{Name: "reward"})
		customersCollection.Fields.Add(&core.JSONField{Name: "reward_history"})
		customersCollection.Fields.Add(&core.TextField{Name: "model"}) // qr code

		customersCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		customersCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		customersCollection.AddIndex("idx_customers_email", true, "email", "")
		customersCollection.AddIndex("idx_customers_model", false, "model", "")

		if err := app.Save(customersCollection); err != nil {
			return err
		}

		log.Println("--- Migration: customers- Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
