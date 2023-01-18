package user_module

import (
	"backend/lib/middleware"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

const (
	GetUsersUrl               = "/user"
	GetUserByIDUrl            = "/user/id/{id}"
	RegisterUserUrl           = "/user/register"
	DeleteUserUrl             = "/user"
	GetUserByUsernameUrl      = "/user/username/{username}"
	LoginUserUrl              = "/user/login"
	GetUserDetailsWithAuthUrl = "/user/details"
)

func Router(api *chi.Mux) {
	api.Get(GetUsersUrl, GetUsers)
	api.Get(GetUserByIDUrl, GetUserByID)
	api.Delete(DeleteUserUrl, DeleteUser)
	api.Get(GetUserByUsernameUrl, GetUserByUsername)
	api.Post(LoginUserUrl, LoginUser)

	// Use authentication Guard for the route that returns non-public
	// user details
	api.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthenticated)
		r.Get(GetUserDetailsWithAuthUrl, GetUserDetailsWithAuth)
	})

	// Use rate limiting on registration route to prevent bot spamming
	// Example -> https://go-chi.io/#/pages/middleware?id=http-rate-limiting-middleware
	api.Group(func(r chi.Router) {
		r.Use(httprate.LimitByIP(100, 100*time.Minute))
		r.Post(RegisterUserUrl, RegisterUser)
	})
}
