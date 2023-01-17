package auth

// Return true only when the input JWT string is valid.
// Note that this returns false, if the token is valid,
// but has expired.
func JwtIsValid(tokenString string) bool {

	token, err := ConvertJwtStringToToken(tokenString)
	if err != nil {
		return false
	}
	if token.Valid {
		return true
	}
	return false
}
