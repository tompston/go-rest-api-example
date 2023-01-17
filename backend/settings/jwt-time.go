package settings

import (
	"time"
)

// Pass in the number of minutes, to convert it to the corresponding time.Time value
func JwtExpirationTime(jwt_token_duration int) time.Time {
	duration := time.Duration(jwt_token_duration)
	expiration_time := time.Now().Add(duration * time.Minute)
	return expiration_time
}

// Return the access token duration that is used in the jwt auth string
func AccessTokenDuration() time.Time {
	return JwtExpirationTime(JWT_ACCESS_TOKEN_DURATION)
}

// Return the refresh token duration that is used in the jwt auth string
func RefreshTokenDuration() time.Time {
	return JwtExpirationTime(JWT_REFRESH_TOKEN_DURATION)
}
