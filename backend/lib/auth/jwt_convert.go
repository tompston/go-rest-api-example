package auth

import (
	"backend/settings"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// If an error occurs while parsing the tokenString, return error
func ConvertJwtStringToToken(tokenString string) (*jwt.Token, error) {

	var JWT_KEY = []byte(settings.JWT_KEY)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error occured while parsing the JWT token")
		}
		return JWT_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	return token, err
}
