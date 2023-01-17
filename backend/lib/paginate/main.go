package paginate

import (
	"backend/settings"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Credits
// https://medium.easyread.co/how-to-do-pagination-in-postgres-with-golang-in-4-common-ways-12365b9fb528
func DecodeCursor(encodedCursor string) (timestamp time.Time, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	timestamp, err = time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return
	}
	uuid = arrStr[1]
	return
}

func EncodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

/*
Pagination param variables that can be used
for querying the data
  - limit   = How many records can be returned
  - key 	= the name of the key which will hold the cursor value
  - cursor 	= the value of the cursor key
  - url		= Full API path of the controller (excluding query params)
*/
func PaginationParams(r *http.Request, url_pagination_key, endpoint_path string) (int32, string, string, string) {

	limit := int32(settings.PAGINATION_LIMIT)
	key := url_pagination_key
	cursor := r.URL.Query().Get(url_pagination_key)
	url := fmt.Sprintf("%s%s", settings.BASE_URL, endpoint_path)

	return limit, key, cursor, url
}

// Return a link that has an appended key-value pair
func LinkWithQueryKey(endpoint_path, url_key_name, url_key_value string) string {
	return endpoint_path + "?" + url_key_name + "=" + url_key_value
}
