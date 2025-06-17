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
		// @request.auth.collectionName = "users"
		log.Println("Creating collection: options")
		optionsCollection := core.NewBaseCollection(constants.TableStoreOptions)

		optionsCollection.ViewRule = types.Pointer("")
		optionsCollection.ListRule = types.Pointer("")
		optionsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")
		optionsCollection.UpdateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")

		optionsCollection.Fields.Add(&core.NumberField{Name: "option_id", Required: true, Min: types.Pointer(0.0)})
		optionsCollection.Fields.Add(&core.TextField{Name: "name"})
		optionsCollection.Fields.Add(&core.TextField{Name: "short_name"})
		optionsCollection.Fields.Add(&core.TextField{Name: "type"})
		optionsCollection.Fields.Add(&core.NumberField{Name: "sort_order"})
		optionsCollection.Fields.Add(&core.NumberField{Name: "title"})
		optionsCollection.Fields.Add(&core.NumberField{Name: "option_group"})
		optionsCollection.Fields.Add(&core.NumberField{Name: "required"})
		optionsCollection.Fields.Add(&core.NumberField{Name: "free_option_count"})
		optionsCollection.Fields.Add(&core.NumberField{Name: "max_option_count"})
		optionsCollection.Fields.Add(&core.JSONField{Name: "option_desc"})
		optionsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		optionsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		optionsCollection.AddIndex("idx_option_id", true, "option_id", "")

		if err := app.Save(optionsCollection); err != nil {
			return err
		}

		// ===========

		log.Println("Creating collection: option_values")

		optionValuesCollection := core.NewBaseCollection(constants.TableStoreOptionValues)
		optionValuesCollection.ViewRule = types.Pointer("")
		optionValuesCollection.ListRule = types.Pointer("")
		optionValuesCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")
		optionValuesCollection.UpdateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")

		// RELATION
		optionValuesCollection.Fields.Add(&core.RelationField{Name: "option", CollectionId: optionsCollection.Id, Required: true, MaxSelect: 1, MinSelect: 1, CascadeDelete: true})
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "option_id", Required: true, Min: types.Pointer(0.0), OnlyInt: true})
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "option_value_id", Required: true, Min: types.Pointer(0.0)})
		optionValuesCollection.Fields.Add(&core.TextField{Name: "name"})
		optionValuesCollection.Fields.Add(&core.TextField{Name: "short_name"})
		optionValuesCollection.Fields.Add(&core.JSONField{Name: "desc"})

		optionValuesCollection.Fields.Add(&core.NumberField{Name: "sort_order"})
		optionValuesCollection.Fields.Add(&core.JSONField{Name: "related_option_id"})

		optionValuesCollection.Fields.Add(&core.NumberField{Name: "price_ht", Required: true, Min: types.Pointer(0.0)})
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "price_ttc", Required: true, Min: types.Pointer(0.0)})

		taxesCollection, err := app.FindCollectionByNameOrId(constants.TableTaxRates)
		if err != nil {
			return err
		}

		optionValuesCollection.Fields.Add(&core.RelationField{Name: "tax_rate", CollectionId: taxesCollection.Id, Required: true, MaxSelect: 1, MinSelect: 1})
		optionValuesCollection.Fields.Add(&core.TextField{Name: "price_status"})
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "reset"})
		optionValuesCollection.Fields.Add(&core.TextField{Name: "grup"})
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "status"}) // is active
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "not_recommanded"})
		optionValuesCollection.Fields.Add(&core.NumberField{Name: "image_type"})
		optionValuesCollection.Fields.Add(&core.TextField{Name: "image_url"})

		optionValuesCollection.Fields.Add(&core.TextField{Name: "model"}) // barcode
		optionValuesCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		optionValuesCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		//indexs
		optionValuesCollection.AddIndex("idx_option_value_id", true, "option_value_id", "")
		optionValuesCollection.AddIndex("idx_store_option", false, "option", "")

		if err := app.Save(optionValuesCollection); err != nil {
			return err
		}

		log.Println("--- Migration: options, option_values (create)  -- Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
