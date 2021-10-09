package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

func CheckRoleAdmin(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return token, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role_id"] != "1" {
			return token, errors.New("Admin role access only !")
		}
	} else {
		return token, errors.New("Couldn't handle this token !")
	}

	return token, err
}
