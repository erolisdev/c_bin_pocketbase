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
		log.Println("Creating collection: products")
		productsCollection := core.NewBaseCollection(constants.TableStoreProducts)

		productsCollection.ViewRule = types.Pointer("")
		productsCollection.ListRule = types.Pointer("")
		productsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")
		productsCollection.UpdateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")

		productsCollection.Fields.Add(&core.NumberField{Name: "product_id", Required: true, Min: types.Pointer(0.0)})
		categoriesCollection, err := app.FindCollectionByNameOrId(constants.TableStoreCategories)
		if err != nil {
			return err
		}

		productsCollection.Fields.Add(&core.RelationField{
			Name:         "category",
			CollectionId: categoriesCollection.Id,
		})

		productsCollection.Fields.Add(&core.NumberField{Name: "category_id", Required: true, Min: types.Pointer(0.0)})
		productsCollection.Fields.Add(&core.TextField{Name: "name", Required: true})
		productsCollection.Fields.Add(&core.TextField{Name: "short_name"})

		// price
		productsCollection.Fields.Add(&core.NumberField{Name: "price_ht", Required: true, Min: types.Pointer(0.0)})
		productsCollection.Fields.Add(&core.NumberField{Name: "price_ttc", Required: true, Min: types.Pointer(0.0)})
		// taxes
		taxesCollection, err := app.FindCollectionByNameOrId(constants.TableTaxRates)
		if err != nil {
			return err
		}
		productsCollection.Fields.Add(&core.RelationField{
			Name:         "tax_rate",
			Required:     true,
			CollectionId: taxesCollection.Id,
			MaxSelect:    1,
			MinSelect:    1,
		})

		productsCollection.Fields.Add(&core.NumberField{Name: "status"}) // is active
		productsCollection.Fields.Add(&core.NumberField{Name: "private"})
		productsCollection.Fields.Add(&core.NumberField{Name: "show_in_suggested"})
		productsCollection.Fields.Add(&core.NumberField{Name: "shipping_code"})
		productsCollection.Fields.Add(&core.NumberField{Name: "isFullWidth"})
		productsCollection.Fields.Add(&core.NumberField{Name: "sort_order"})
		productsCollection.Fields.Add(&core.NumberField{Name: "c_sort_order"})
		productsCollection.Fields.Add(&core.TextField{Name: "image_url"})
		productsCollection.Fields.Add(&core.JSONField{Name: "related"})
		productsCollection.Fields.Add(&core.JSONField{Name: "description"})
		//
		productsCollection.Fields.Add(&core.NumberField{Name: "printer_id"})
		productsCollection.Fields.Add(&core.NumberField{Name: "template_id", Required: true})
		productsCollection.Fields.Add(&core.TextField{Name: "label_printer_id"}) // modified

		// productsCollection.Fields.Add(&core.TextField{Name: "printer_ip", Required: true})
		// productsCollection.Fields.Add(&core.NumberField{Name: "printer_port"})
		productsCollection.Fields.Add(&core.TextField{Name: "model"}) // barcode
		productsCollection.Fields.Add(&core.JSONField{Name: "product_option_ids"})

		// categories
		// categoryCollection, err := app.FindCollectionByNameOrId(constants.TableStoreCategories)
		// if err != nil {
		// 	return err
		// }

		// productsCollection.Fields.Add(&core.RelationField{Name: "category", CollectionId: categoryCollection.Id})

		// options
		// optionsCollection, err := app.FindCollectionByNameOrId(constants.TableStoreOptions)
		// if err != nil {
		// 	return err
		// }

		// productsCollection.Fields.Add(&core.RelationField{Name: "category", CollectionId: optionsCollection.Id, CascadeDelete: false})

		productsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		productsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		// indexs
		productsCollection.AddIndex("idx_product_id", true, "product_id", "")
		productsCollection.AddIndex("idx_category_id", false, "category", "")

		if err := app.Save(productsCollection); err != nil {
			return err
		}

		log.Println("--- Migration: create_products (Apply - Güncel API) - Tamamlandı ---")
		return nil
	}, func(app core.App) error {
		// Revert
		return nil
	})
}
