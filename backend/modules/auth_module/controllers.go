package auth_module

import (
	"backend/lib/auth"
	res "backend/lib/response"

	"net/http"

	"github.com/go-chi/chi/v5"

	"backend/lib/validate"
	"encoding/json"
)

const (
	TokenIsValidUrl = "/auth/token"
)

func Router(api *chi.Mux) {
	api.Post(TokenIsValidUrl, TokenIsValid)
}

func TokenIsValid(w http.ResponseWriter, r *http.Request) {

	// validate the sent json object
	payload := new(TokenIsValidBody)

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		res.Response(w, 400, err.Error(), res.FailedJsonValidation)
		return
	}
	if err := validate.ValidateStruct(payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidation)
		return
	}

	token_is_valid := auth.JwtIsValid(payload.AccessToken)
	if !token_is_valid {
		res.Response(w, 401, nil, "Authorization token is invalid!")
		return
	}

	res.Response(w, 200, token_is_valid, "")
}
