package app

import (
	res "backend/lib/response"
	"backend/settings"

	"net/http"

	"github.com/go-chi/chi/v5"

	"backend/modules/auth_module"
	"backend/modules/transaction_module"
	"backend/modules/user_module"
)

func Router(app *chi.Mux) {

	// home_view
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res.Response(w, 200, nil, "Hello There!")
	})

	api := chi.NewRouter()
	app.Mount(settings.API_PATH, api)

	user_module.Router(api)
	transaction_module.Router(api)
	auth_module.Router(api)
}
