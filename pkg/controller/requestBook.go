package controller

import (
	"fmt"
	"lms/pkg/models"
	"lms/pkg/utils"
	"net/http"
	"strconv"
)

func RequestBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		BID := r.FormValue("checkInID")
		BIDint, _ := strconv.Atoi(BID)
		fmt.Println(BID)
		cookie, err := r.Cookie("access-token")
		if err != nil {
			fmt.Println("Cookie Does Not exist")
		} else {
			token := cookie.Value
			claims, err := utils.DecodeJWT(token)
			if err != nil {
				fmt.Println("Error in decoding")
			} else {
				resp := models.RequestBook(claims["UID"].(int), BIDint)
				if resp != "" {
					fmt.Println(resp)
				}
				http.Redirect(w, r, "/browse", http.StatusSeeOther)
				return

			}
		}

	}
}
