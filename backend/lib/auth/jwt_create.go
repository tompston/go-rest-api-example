package auth

import (
	"backend/settings"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define which fields should be stored in the claims.
// THe json field will be used as the key name for the
// claim, when decoded
type TokenClaims struct {
	*jwt.RegisteredClaims
	IsAdmin   bool   `json:"is_admin"`
	Username  string `json:"username"`
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"exp"`
}

// https://www.golinuxcloud.com/golang-jwt/
// Create a JWT string that stores the passed in fields.
// If at some point the creation of the token fails,
// return an empty string and an error
func CreateJwtToken(user_claims TokenClaims, expiration_time time.Time) (string, error) {

	// token := jwt.New(jwt.SigningMethodES256)
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &TokenClaims{
		IsAdmin:   user_claims.IsAdmin,
		Username:  user_claims.Username,
		UserID:    user_claims.UserID,
		ExpiresAt: expiration_time.Unix()}

	tokenString, err := token.SignedString([]byte(settings.JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// Create a JWT Access Token string, using the expiration time that is used in the settings
func CreateJwtAccesToken(user_claims TokenClaims) (string, error) {
	token, err := CreateJwtToken(user_claims, settings.AccessTokenDuration())
	if err != nil {
		return "", err
	}
	return token, err
}

// Create a JWT Refresh Token string, using the expiration time that is used in the settings
func CreateJwtRefreshToken(user_claims TokenClaims) (string, error) {
	token, err := CreateJwtToken(user_claims, settings.RefreshTokenDuration())
	if err != nil {
		return "", err
	}
	return token, err
}

// Return claims from the token, if no errors or token is valid
func GetTokenClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := ConvertJwtStringToToken(tokenString)
	if err != nil {
		fmt.Println("Could not convert token string to JWT!")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("jwt token is invalid")
	}

	return claims, err
}
