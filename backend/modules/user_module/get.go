package user_module

import (
	"backend/db/sqlc"
	"backend/lib/auth"
	"backend/lib/paginate"
	"backend/lib/request"
	res "backend/lib/response"
	"backend/settings/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Return all users using cursor pagination.
//   - If the query param is an empty string, execute
//     the first pagination func which does not need any state params.
//   - If the query param exists and is valid, exec the query that
//     does cursor pagination.
func GetUsers(w http.ResponseWriter, r *http.Request) {

	limit, url_pagination_key, cursor, endpoint_path := paginate.
		PaginationParams(r, "cursor", GetUsersUrl)

	ctx := context.Background()

	// response if there is no cursor query param
	if cursor == "" {
		data, err := sqlc.New(database.DB).User_GetAllWithPaginationFirstPage(ctx, limit+1)
		if err != nil {
			res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
			return
		}

		next_cursor := paginate.EncodeCursor(
			data[limit].CreatedAt,
			data[limit].UserID.String())

		res.ResponseWithPagination(w, 200, data[:limit], "",
			res.PaginationLinks{Next: endpoint_path + "?" + url_pagination_key + "=" + next_cursor})
		return
	}

	// response if cursor query param exists and is valid
	if cursor != "" {
		timestamp, id, err := paginate.DecodeCursor(cursor)
		if err != nil {
			res.Response(w, 400, nil, "Could not decode the cursor!")
			return
		}

		data, err := sqlc.New(database.DB).
			User_GetAllWithPaginationNextPage(ctx,
				sqlc.User_GetAllWithPaginationNextPageParams{
					UserID:    uuid.MustParse(id),
					CreatedAt: timestamp,
					Limit:     limit + 1})

		if err != nil {
			res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
			return
		}

		// If lenght of the returned array does not have the last
		// row, you can't create the next cursor
		if len(data) < int(limit) {
			res.ResponseWithPagination(w, 200, data, "", res.PaginationLinks{Next: "null"})
			return
		}

		next_cursor := paginate.EncodeCursor(
			data[limit].CreatedAt,
			data[limit].UserID.String())

		res.ResponseWithPagination(w, 200, data[:limit], "", res.PaginationLinks{
			Next: endpoint_path + "?" + url_pagination_key + "=" + next_cursor})
		return
	}
}

// If the user_id exists, return response with public user details
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	user_id, err := uuid.Parse(id)
	if err != nil {
		res.Response(w, 400, nil, "Could not parse the user id!")
		return
	}

	data, err := sqlc.New(database.DB).User_GetWhereIdEquals(
		context.Background(), user_id)

	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	res.Response(w, 200, data, "User Found!")
}

// If the username exists, return response with public user details
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	username := chi.URLParam(r, "username")

	data, err := sqlc.New(database.DB).
		User_GetWhereUsernameEquals(context.Background(), username)

	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	// if the returned data from the db is empty, the user
	// with the specified username does not exist
	if (data == sqlc.User_GetWhereUsernameEqualsRow{}) {
		res.Response(w, 400, nil, "User does not exist!")
		return
	}

	res.Response(w, 200, data, "")
}

func GetUserDetailsWithAuth(w http.ResponseWriter, r *http.Request) {

	auth_header := request.GetAuthHeader(r)
	auth_claim, err := auth.GetTokenClaims(auth_header)
	if err != nil {
		res.Response(w, 401, nil, "Authentication Failed! Invalid token claim extraction.")
		return
	}

	user_id := uuid.MustParse(auth_claim["user_id"].(string))
	ctx := context.Background()
	db := sqlc.New(database.DB)

	data, err := db.User_GetWhereIdEquals(ctx, user_id)
	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	// If no transactions are found for the user, the returned balance is 0
	balance, err := db.Balance_GetUserBalanceByUserID(ctx, user_id)
	if err != nil {
		fmt.Println("Error occured during calculating the balance of the user, because no transactions could be found, user_id :: ", user_id)
	}

	// Combine user data and balance into a single struct
	type Response struct {
		UserID    uuid.UUID `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		Username  string    `json:"username"`
		Balance   int64     `json:"balance"`
	}

	full_data := Response{
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		Username:  data.Username,
		Balance:   balance,
	}

	res.Response(w, 200, full_data, "")
}
