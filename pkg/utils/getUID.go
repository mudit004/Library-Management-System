package utils

import (
	"fmt"
	"net/http"
)

func GetUID(w http.ResponseWriter, r *http.Request) interface{} {
	if r.Method == http.MethodGet {
		cookie, err := r.Cookie("access-token")
		if err != nil {
			fmt.Println("Cookie Does Not Exist")
			return nil
		} else {
			token := cookie.Value
			claims, err := DecodeJWT(token)
			if err != nil {
				fmt.Println("Error in decoding")
				return nil
			} else {
				UID := claims["UID"]
				fmt.Println(UID)
				return UID
			}
		}
	}
	return nil
}
