package transaction_module

import (
	"backend/db/sqlc"
	"backend/lib/auth"
	"backend/lib/paginate"
	"backend/lib/request"
	res "backend/lib/response"
	"backend/lib/utils"
	"backend/settings/database"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Return all of the sent transactions for the logged in user
// using cursor paginations
func GetTransactionsForUser(w http.ResponseWriter, r *http.Request) {

	auth_header := request.GetAuthHeader(r)
	auth_claim, err := auth.GetTokenClaims(auth_header)
	if err != nil {
		res.Response(w, 401, nil, "Authentication Failed!")
	}

	user_id := uuid.MustParse(auth_claim["user_id"].(string))

	limit, url_pagination_key, cursor, endpoint_path := paginate.PaginationParams(r, "cursor", GetTransactionsForUserUrl)
	ctx := context.Background()

	// If no cursor provided, execute the first query
	if cursor == "" {
		data, err := sqlc.New(database.DB).
			User_GetSentTransactionsWhereUserIdEqualsFirstPage(ctx,
				sqlc.User_GetSentTransactionsWhereUserIdEqualsFirstPageParams{
					UserID: user_id, Limit: limit + 1})

		if err != nil {
			res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
			return
		}

		// If no transactions for the user are found, return nil data,
		// and an appropriate message.
		if data == nil {
			res.Response(w, 200, nil, "No transactions for user found!")
			return
		}

		next_cursor := paginate.EncodeCursor(
			data[limit].CreatedAt,
			data[limit].UserID.String())

		res.ResponseWithPagination(w, 200, data[:limit], "",
			res.PaginationLinks{
				Next: endpoint_path + "?" + url_pagination_key + "=" + next_cursor,
			})
		return
	}

	// If cursor exists and is valid, execute the pagination query
	if cursor != "" {
		timestamp, _, err := paginate.DecodeCursor(cursor)
		if err != nil {
			res.Response(w, 400, nil, "Could not decode the cursor!")
			return
		}

		data, err := sqlc.New(database.DB).
			User_GetSentTransactionsWhereUserIdEquals(
				ctx, sqlc.User_GetSentTransactionsWhereUserIdEqualsParams{
					UserID:    user_id,
					CreatedAt: timestamp,
					Limit:     limit + 1})

		if err != nil {
			res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
			return
		}

		if len(data) < int(limit) {
			res.ResponseWithPagination(w, 200, data, "", res.PaginationLinks{Next: "null"})
			return
		}

		next_cursor := paginate.EncodeCursor(
			data[limit].CreatedAt,
			data[limit].UserID.String())

		res.ResponseWithPagination(w, 200, data[:limit], "",
			res.PaginationLinks{
				Next: endpoint_path + "?" + url_pagination_key + "=" + next_cursor,
			})
		return
	}
}

// Not implemented
func GetTransactionByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	_ = id

	// replace the function to return the data from the db
	data, err := utils.PlaceholderFunction()
	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	res.Response(w, 200, data, "")
}
