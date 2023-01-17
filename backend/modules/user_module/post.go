package user_module

import (
	"backend/db/sqlc"
	"backend/lib/auth"
	res "backend/lib/response"
	"backend/lib/validate"
	"backend/settings/database"
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// validate the sent json object
	payload := new(RegisterUserBody)
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		res.Response(w, 400, err.Error(), res.FailedJsonValidation)
		return
	}
	if err := validate.ValidateStruct(payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidation)
		return
	}

	hashed_password, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		res.Response(w, 400, nil, "Could not hash the password!")
		return
	}

	// replace the function to return the data from the db
	data, err := sqlc.New(database.DB).User_Create(
		context.Background(),
		sqlc.User_CreateParams{
			Username: payload.Username,
			Password: string(hashed_password),
			Email:    payload.Email})

	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	res.Response(w, 201, data, "New User created!")
}

// Return access and refresh JWT tokens in the response, if
// the sent credentials are correct
func LoginUser(w http.ResponseWriter, r *http.Request) {

	// validate the sent json object
	payload := new(LoginUserBody)
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		res.Response(w, 400, err.Error(), res.FailedJsonValidation)
		return
	}
	if err := validate.ValidateStruct(payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidation)
		return
	}

	data, err := sqlc.New(database.DB).User_LoginWithUsername(
		context.Background(), payload.Username)

	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	// If no user is returned from the db, return error
	if (data == sqlc.User{}) {
		res.Response(w, 400, nil, "User does not exist!")
		return
	}

	// If the user exists and the payload password does
	// not match the decoded hashed password in the db,
	// return an error message and StatusUnauthorized code
	if err := bcrypt.CompareHashAndPassword(
		[]byte(data.Password),
		[]byte(payload.Password)); err != nil {
		res.Response(w, 401, nil, "Incorrect password provided!")
		return
	}

	token_claims := auth.TokenClaims{
		IsAdmin:  data.IsAdmin,
		Username: data.Username,
		UserID:   data.UserID.String()}

	access_token, err := auth.CreateJwtAccesToken(token_claims)
	if err != nil || access_token == "" {
		res.Response(w, 400, nil, "Could not create a valid JWT Access Token!")
		return
	}

	refresh_token, err := auth.CreateJwtRefreshToken(token_claims)
	if err != nil || refresh_token == "" {
		res.Response(w, 400, nil, "Could not create a valid JWT Refresh Token!")
		return
	}

	data_with_token := &res.ResponseWithJwt{
		Data: data,
		JwtFields: res.JwtFields{
			AccessToken:  access_token,
			RefreshToken: refresh_token}}

	// fmt.Println(access_token)
	// fmt.Println(refresh_token)

	res.Response(w, 200, data_with_token, "User Logged in!")
}
