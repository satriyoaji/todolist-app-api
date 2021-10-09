package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
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

func MakeExpiredToken(r *http.Request) error {
	tokenString := ExtractToken(r)

	//claims := jwt.MapClaims{}
	//_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	//	fmt.Println("claims awal: ", claims)
	//	claims["exp"] = time.Now().Add(time.Second * 10).Unix()
	//	fmt.Println("claims awal: ", claims)
	//	return []byte("<YOUR VERIFICATION KEY>"), nil
	//})

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = time.Now().Add(time.Second * 1).Unix()
	} else {
		return errors.New("Couldn't handle this token !")
	}

	return err
}
