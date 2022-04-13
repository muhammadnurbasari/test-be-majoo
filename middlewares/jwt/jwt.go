package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	applicationName  = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("APP_SECRET_KEY_JWT"))
	loginExpDuration = time.Duration(60*12) * time.Minute // 12 hour
)

//MyClaims - struct jwt
type MyClaims struct {
	jwt.StandardClaims
	ID int64 `json:"id"`
}

//GenerateTokenJwt - generate token jwt
func GenerateTokenJwt(id int64) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: time.Now().Add(loginExpDuration).Unix(),
		},
		ID: id,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", errors.New("GenerateTokenJwt err = " + err.Error())
	}

	return signedToken, nil
}
