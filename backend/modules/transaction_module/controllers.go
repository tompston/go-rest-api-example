package transaction_module

import (
	"backend/lib/middleware"

	"github.com/go-chi/chi/v5"
)

const (
	GetTransactionsForUserUrl = "/transaction"
	CreateTransactionUrl      = "/transaction"
	GetTransactionByIDUrl     = "/transaction/{transaction_id}"
)

func Router(api *chi.Mux) {
	// If user is authenticated, allow access to these routes
	api.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthenticated)
		r.Get(GetTransactionsForUserUrl, GetTransactionsForUser)
		r.Get(GetTransactionByIDUrl, GetTransactionByID)
		r.Post(CreateTransactionUrl, CreateTransaction)
	})
}
