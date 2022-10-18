package authorizationhandler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("dummySignedKey")

type authHander struct {
	l *log.Logger
}

func NewAuthHander(l *log.Logger) *authHander {
	return &authHander{l}
}

func (authHander *authHander) GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "@binator_1308"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		//fmt.Errorf("generating JWT Token failed %v",err)
		return "", err
	}
	return tokenString, nil
}

func (authHander *authHander) getToken(w http.ResponseWriter, r *http.Request) {
	if token, err := authHander.GenerateJWTToken(); err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, token)
	}
}

func (authHander *authHander) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
				}
				return mySigningKey, nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				authHander.l.Println(claims["user"], claims["exp"])
				next.ServeHTTP(w, r)
			} else {
				fmt.Fprintf(w, err.Error())
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
