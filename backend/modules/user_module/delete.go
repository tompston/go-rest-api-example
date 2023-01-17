package user_module

import (
	res "backend/lib/response"
	"backend/lib/utils"
	"net/http"
)

// Not implemented
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// replace the function to return the data from the db
	data, err := utils.PlaceholderFunction()
	if err != nil {
		res.Response(w, 400, nil, res.DbErrorMessage(err.Error()))
		return
	}

	res.Response(w, 200, data, "")
}
