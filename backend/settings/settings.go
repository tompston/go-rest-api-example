package settings

import (
	"fmt"
)

// .env variables. Name of the variable is mapped to the env var name.
var (
	// Go .env vars
	GO_BACKEND_PORT = Config("GO_BACKEND_PORT")
	DEBUG_MODE      = Config("DEBUG_MODE")
	HOST_URL        = Config("HOST_URL")
	API_PATH        = Config("API_PATH")
	BASE_URL        = fmt.Sprintf("%s%s", HOST_URL, API_PATH) // host + api prefix path
	// DB .env vars
	DB_PASS = Config("DB_PASS")
	DB_HOST = Config("DB_HOST")
	DB_USER = Config("DB_USER")
	DB_NAME = Config("DB_NAME")
	DB_PORT = Config("DB_PORT")
	DB_SSL  = Config("DB_SSL")
	DB_TZ   = Config("DB_TZ")

	// JWT Variables
	JWT_KEY = Config("JWT_KEY")
)

const (
	PAGINATION_LIMIT           = 20 // specify how many rows can be returned in a paginated response
	JWT_ACCESS_COOKIE_NAME     = "access_cookie"
	JWT_REFRESH_COOKIE_NAME    = "refresh_cookie"
	JWT_ACCESS_TOKEN_DURATION  = 10 // time in minutes
	JWT_REFRESH_TOKEN_DURATION = 15 // time in minutes
	//
	AUTHENTICATION_HEADER_KEY = "Authorization"
)
