package utils

import (
	"fmt"
	"net/http"
)

func GetUserID(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method == http.MethodGet {
		cookie, err := r.Cookie("access-token")
		if err != nil {
			return nil, err
		} else {
			token := cookie.Value
			claims, err := DecodeJWT(token)
			if err != nil {
				return nil, err
			} else {
				UserID := claims["UserID"]
				return UserID, nil
			}
		}
	}
	return nil, fmt.Errorf("request not found")
}
