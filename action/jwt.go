package action

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const SECRET_JWT = "Thisissecret"

func Create(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token expired after 12 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}

func Extract(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	return strings.Split(bearerToken, " ")[1]
}

func Verify(r *http.Request) error {
	tokenString := Extract(r)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_JWT), nil
	})

	if err != nil {
		return err
	}
	return nil
}
