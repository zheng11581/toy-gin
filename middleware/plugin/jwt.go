package plugin

import (
	"github.com/golang-jwt/jwt/v4"
)

var key = "abcdefg1234567"

type Data struct {
	Name   string
	Age    int
	Gender int
	jwt.RegisteredClaims
}

func Sign(data jwt.Claims) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Sign and get the complete encoded token as a string using the secret
	sign, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return sign, nil
}

func Verify(sign string, data jwt.Claims) error {

	_, err := jwt.ParseWithClaims(sign, data, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return err

}
