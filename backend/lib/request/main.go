package request

import (
	"backend/settings"
	"net/http"
	"strings"
)

func GetAuthHeader(r *http.Request) string {
	auth_key := r.Header.Get(settings.AUTHENTICATION_HEADER_KEY)
	auth_token := strings.Split(auth_key, "Bearer ")[1]
	return auth_token
}
