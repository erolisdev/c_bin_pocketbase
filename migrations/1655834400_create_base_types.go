package migrations

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations" // Alias for migrations package
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		// --- taxes ---
		log.Println("Colection creating: taxes")

		taxesCollection := core.NewBaseCollection("tax_rates")
		taxesCollection.Fields.Add(&core.TextField{Name: "description"})
		taxesCollection.Fields.Add(&core.TextField{Name: "type", Required: true})

		taxesCollection.Fields.Add(&core.NumberField{
			Name:     "rate",
			Required: true,
			Min:      types.Pointer(0.0),
			Max:      types.Pointer(100.0), // Oran 0-100 arası
		})

		taxesCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		taxesCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		if err := app.Save(taxesCollection); err != nil {
			return err
		}
		// --- default tax rates ---
		var franceTaxRates = []map[string]any{
			{
				"type":        "D",
				"description": "SuperReduced",
				"rate":        2.1,
			},
			{
				"type":        "C",
				"description": "Reduced",
				"rate":        5.5,
			},
			{
				"type":        "B",
				"description": "Intermediate",
				"rate":        10.0,
			},
			{
				"type":        "A",
				"description": "Standard",
				"rate":        20.0,
			},
		}

		//save default taxe list
		for _, tax := range franceTaxRates {
			record := core.NewRecord(taxesCollection)
			for key, value := range tax {
				record.Set(key, value)
			}

			if err := app.Save(record); err != nil {
				return err
			}
		}

		// --- categories ---
		log.Println("Colection creating: categories")
		categoriesCollection := core.NewBaseCollection("categories")
		categoriesCollection.Fields.Add(&core.TextField{Name: "name", Required: true})
		categoriesCollection.Fields.Add(&core.NumberField{Name: "category_id"})
		categoriesCollection.Fields.Add(&core.NumberField{Name: "column"})
		categoriesCollection.Fields.Add(&core.NumberField{Name: "status"})
		categoriesCollection.Fields.Add(&core.NumberField{Name: "sort_order"})
		categoriesCollection.Fields.Add(&core.NumberField{Name: "show_in_suggested"})
		categoriesCollection.Fields.Add(&core.TextField{Name: "image_url"})
		categoriesCollection.Fields.Add(&core.JSONField{Name: "description"})
		categoriesCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		categoriesCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		if err := app.Save(categoriesCollection); err != nil {
			return err
		}

		// --- payment_methods ---
		log.Println("Colection creating: payment_methods")
		paymentMethodsCollection := core.NewBaseCollection("payment_methods")
		paymentMethodsCollection.Fields.Add(&core.TextField{Name: "name", Required: true})
		paymentMethodsCollection.Fields.Add(&core.TextField{Name: "type", Required: true})
		paymentMethodsCollection.Fields.Add(&core.BoolField{Name: "status"})
		paymentMethodsCollection.Fields.Add(&core.AutodateField{Name: "updated", OnUpdate: true})
		paymentMethodsCollection.Fields.Add(&core.AutodateField{Name: "created", OnCreate: true})

		if err := app.Save(paymentMethodsCollection); err != nil {
			return err
		}

		var francePaymentMethods = []map[string]any{
			{
				"type":   "STRIPE",
				"status": true,
				"name":   "Stripe",
			},
			{
				"type":   "TRANSFER",
				"status": true,
				"name":   "Virement bancaire",
			},
			{
				"type":   "CHECK",
				"status": true,
				"name":   "Chèque",
			},
			{
				"type":   "VOUCHER",
				"status": true,
				"name":   "Bon d'achat",
			},
			{
				"type":   "TICKET RESTAURANT CARD",
				"status": true,
				"name":   "Carte Ticket Restaurant",
			},
			{
				"type":   "TICKET RESTAURANT",
				"status": true,
				"name":   "Ticket restaurant",
			},
			{
				"type":   "CARD",
				"status": true,
				"name":   "Carte bancaire",
			},
			{
				"type":   "CASH",
				"status": true,
				"name":   "Espèces",
			},
		}

		//save default taxe list
		for _, tax := range francePaymentMethods {
			record := core.NewRecord(paymentMethodsCollection)
			for key, value := range tax {
				record.Set(key, value)
			}

			if err := app.Save(record); err != nil {
				return err
			}
		}

		log.Println("--- Migration: create_base_types - Tamamlandı ---")
		return nil
	}, nil)
}
