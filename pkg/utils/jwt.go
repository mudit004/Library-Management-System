package utils

import (
	"fmt"
	"lms/pkg/types"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/slices"
)

func NewJWT(user types.User) string {
	config, err := NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	key := []byte(config.JWT_SECRET_KEY)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = user.UName
	claims["UID"] = user.UID
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	tokenStr, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return tokenStr

}

func DecodeJWT(tokenStr string) (jwt.MapClaims, error) {
	config, err := NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	key := []byte(config.JWT_SECRET_KEY)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Token claims could not be extracted")
	}

	return claims, nil

}

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goThroughUrls := []string{"/", "/register", "/login", "/registerUser", "/admin", "/adminLogin", "/internalServerError"}
		if slices.Contains(goThroughUrls, r.URL.Path) || strings.Split(r.URL.Path, "/")[1] == "static" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("access-token")
		if err == nil {
			// Cookie exists
			token := cookie.Value
			claims, err := DecodeJWT(token)
			if err != nil {
				fmt.Println(claims)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

}
