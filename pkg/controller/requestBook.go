package controller

import (
	"lms/pkg/models"
	"lms/pkg/utils"
	"net/http"
	"strconv"
)

func RequestBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		BID := r.FormValue("checkInID")
		BIDint, _ := strconv.Atoi(BID)
		cookie, err := r.Cookie("access-token")
		if err != nil {
			return
		} else {
			token := cookie.Value
			claims, err := utils.DecodeJWT(token)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
				return
			} else {
				_ = models.RequestBook(claims["UID"], BIDint)
				http.Redirect(w, r, "/requestBook", http.StatusSeeOther)
				return

			}
		}

	}
}
