package middleware

import (
	"backend/lib/auth"
	res "backend/lib/response"
	"backend/settings"
	"fmt"
	"net/http"
	"strings"
)

// Get the request header
//   - If authorization header does not exists, return error response
//   - If authorization header exists, but is not valid, return error response
//   - If the auth token is not valid, return error response
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth_key := r.Header.Get(settings.AUTHENTICATION_HEADER_KEY)
		if auth_key == "" {
			res.Response(w, 401, nil, "No Authorization header provided!")
			return
		}

		auth_token := strings.Split(auth_key, "Bearer ")[1]
		if auth_token == "" {
			res.Response(w, 401, nil, "No Bearer token provided!")
			return
		}

		auth_token_is_valid := auth.JwtIsValid(auth_token)
		if !auth_token_is_valid {
			res.Response(w, 401, nil, "Invalid Bearer Token provided!")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Get the request header
//   - If authorization header does not exists, return error response
//   - If authorization header exists, but is not valid, return error response
//   - If the auth token is not valid, return error response
//
// TODO : Find out if you can reuse the previous middleware,
// so that the code would not be duplicated.
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth_key := r.Header.Get(settings.AUTHENTICATION_HEADER_KEY)
		if auth_key == "" {
			res.Response(w, 401, nil, "No Authorization header provided!")
			return
		}

		auth_token := strings.Split(auth_key, "Bearer ")[1]
		if auth_token == "" {
			res.Response(w, 401, nil, "No Bearer token provided!")
			return
		}

		auth_token_is_valid := auth.JwtIsValid(auth_token)
		if !auth_token_is_valid {
			res.Response(w, 401, nil, "Invalid Bearer Token provided!")
			return
		}

		auth_token_claims, err := auth.GetTokenClaims(auth_token)
		if err != nil {
			res.Response(w, 401, nil, "Could not extract claims from the authorization token!")
			return
		}

		is_admin := auth_token_claims["is_admin"].(bool)
		if !is_admin {
			res.Response(w, 401, nil, "Could not extract claims from the authorization token!")
			return
		}
		fmt.Println("IS ADMIN :: ", is_admin)

		next.ServeHTTP(w, r)
	})
}
