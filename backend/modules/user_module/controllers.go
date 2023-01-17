package user_module

import (
	"backend/lib/middleware"

	"github.com/go-chi/chi/v5"
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
	api.Post(RegisterUserUrl, RegisterUser)
	api.Delete(DeleteUserUrl, DeleteUser)
	api.Get(GetUserByUsernameUrl, GetUserByUsername)
	api.Post(LoginUserUrl, LoginUser)

	api.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthenticated)
		r.Get(GetUserDetailsWithAuthUrl, GetUserDetailsWithAuth)
	})
}
