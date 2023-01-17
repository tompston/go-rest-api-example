package response

import (
	"encoding/json"
	"net/http"

	"backend/settings"
	"fmt"
)

const (
	ParamIsNotIntMessage    = "Only integers as URL params allowed!"
	FailedDbConnMessage     = "Could not connect to the database!"
	FailedJsonValidation    = "Invalid JSON sent!"
	FailedPayloadValidation = "Payload validation failed!"
)

type Res struct {
	// Baseline Response
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResWithPagination struct {
	// Baseline Response + Pagination links
	Res
	Links interface{} `json:"links"`
}

type PaginationLinks struct {
	Prev string `json:"prev,omitempty"`
	Next string `json:"next,omitempty"`
}

type ResponseWithJwt struct {
	Data      interface{} `json:"data"`
	JwtFields JwtFields   `json:"token,omitempty"`
}

type JwtFields struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// all of the correct status codes can be found here -> https://pkg.go.dev/net/http?utm_source=gopls#StatusOK
func Response(w http.ResponseWriter, status int, data interface{}, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&Res{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func ResponseWithPagination(w http.ResponseWriter, status int, data interface{}, message string, links PaginationLinks) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&ResWithPagination{
		Res: Res{
			Status:  status,
			Message: message,
			Data:    data,
		},
		Links: PaginationLinks{
			Prev: links.Prev,
			Next: links.Next,
		},
	})
}

// return the full error message only during debug
func DbErrorMessage(err_msg string) string {
	if settings.DEBUG_MODE == "true" {
		return fmt.Sprintln(FailedDbConnMessage, err_msg)
	}
	return fmt.Sprintln(FailedDbConnMessage)
}
