package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("X-API-Key")

	return bearerToken
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}