package controller

import (
	"lms/pkg/models"
	"lms/pkg/utils"
	"net/http"
	"strconv"
)

func ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		BookID := r.FormValue("checkInID")
		BookIDint, err := strconv.Atoi(BookID)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}
		cookie, err := r.Cookie("access-token")
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		} else {
			token := cookie.Value
			claims, err := utils.DecodeJWT(token)

			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
				return
			} else {
				err = models.ReturnBook(claims["UID"], BookIDint)
				if err != nil {
					http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
					return
				}
				http.Redirect(w, r, "/requestBook", http.StatusSeeOther)
				return

			}
		}

	}
}
