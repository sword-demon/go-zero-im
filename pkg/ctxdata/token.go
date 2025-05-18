package ctxdata

import "github.com/golang-jwt/jwt"

const Identify = "wujiedeyouxi"

// GetJwtToken 生产jwt 的token
func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
