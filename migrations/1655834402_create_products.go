package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		log.Println("Creating collection: products")
		productsCollection := core.NewBaseCollection("products")

		productsCollection.ViewRule = types.Pointer("")
		productsCollection.ListRule = types.Pointer("")
		productsCollection.CreateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")
		productsCollection.UpdateRule = types.Pointer("@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true")

		productsCollection.Fields.Add(&core.NumberField{Name: "product_id", Required: true, Min: types.Pointer(0.0)})
		productsCollection.Fields.Add(&core.TextField{Name: "name", Required: true})
		productsCollection.Fields.Add(&core.NumberField{Name: "price_ht", Required: true, Min: types.Pointer(0.0)})
		productsCollection.Fields.Add(&core.NumberField{Name: "price_ttc", Required: true, Min: types.Pointer(0.0)})
		productsCollection.Fields.Add(&core.NumberField{Name: "status"}) // is active
		productsCollection.Fields.Add(&core.NumberField{Name: "private"})
		productsCollection.Fields.Add(&core.NumberField{Name: "show_in_suggested"})
		productsCollection.Fields.Add(&core.NumberField{Name: "shipping"})
		productsCollection.Fields.Add(&core.NumberField{Name: "category_id"})

		productsCollection.Fields.Add(&core.NumberField{Name: "isFullWidth"})
		productsCollection.Fields.Add(&core.NumberField{Name: "sort_order"})
		productsCollection.Fields.Add(&core.NumberField{Name: "c_sort_order"})
		productsCollection.Fields.Add(&core.TextField{Name: "image_url"})
		productsCollection.Fields.Add(&core.JSONField{Name: "related"})
		productsCollection.Fields.Add(&core.JSONField{Name: "description"})
		//
		productsCollection.Fields.Add(&core.NumberField{Name: "printer_id"})
		productsCollection.Fields.Add(&core.TextField{Name: "printer_ip", Required: true})
		productsCollection.Fields.Add(&core.NumberField{Name: "printer_port"})
		productsCollection.Fields.Add(&core.NumberField{Name: "template_id", Required: true})
		productsCollection.Fields.Add(&core.TextField{Name: "label_printer_id", Required: true})
		productsCollection.Fields.Add(&core.TextField{Name: "model"}) // barcode

		// İlişkiler eklendi

		// taxes
		taxesCollection, err := app.FindCollectionByNameOrId("tax_rates")
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

		// categories
		categoryCollection, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			return err
		}

		productsCollection.Fields.Add(&core.RelationField{Name: "category", CollectionId: categoryCollection.Id})

		// options
		optionsCollection, err := app.FindCollectionByNameOrId("options")
		if err != nil {
			return err
		}

		productsCollection.Fields.Add(&core.RelationField{Name: "category", CollectionId: optionsCollection.Id, CascadeDelete: false})

		productsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		productsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		// indexs
		productsCollection.AddIndex("idx_product_id", true, "product_id", "")

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
