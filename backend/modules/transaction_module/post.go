package transaction_module

import (
	"backend/db/sqlc"
	res "backend/lib/response"
	"backend/lib/validate"
	"backend/settings/database"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	// validate the sent json object
	payload := new(CreateTransactionBody)
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		res.Response(w, 400, err.Error(), res.FailedJsonValidation)
		return
	}
	if err := validate.ValidateStruct(payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidation)
		return
	}

	// Store the transaction between two accounts in the db
	data, err := sqlc.New(database.DB).Transaction_Create(
		context.Background(),
		sqlc.Transaction_CreateParams{
			SenderID:   uuid.MustParse(payload.SenderId),
			ReceiverID: uuid.MustParse(payload.RecieverId),
			Amount:     payload.Amount})

	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	res.Response(w, 200, data, "")
}
