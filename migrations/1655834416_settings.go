package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {

		settings := app.Settings()

		// for all available settings fields you could check
		// https://github.com/pocketbase/pocketbase/blob/develop/core/settings_model.go#L121-L130
		settings.Meta.AppName = "ERONOR"
		// settings.Meta.HideControls = true
		// settings.Meta.AppURL = "https://example.com"

		settings.Logs.MaxDays = 0
		settings.Logs.LogAuthId = true
		settings.Logs.LogIP = true
		settings.Logs.MinLevel = 4

		return app.Save(settings)
	}, nil)
}
