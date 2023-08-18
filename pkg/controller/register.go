package controller

import (
	"lms/pkg/models"
	"lms/pkg/views"
	"net/http"

	"github.com/tawesoft/golib/v2/dialog"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("Register")
	t.Execute(writer, nil)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		ConfirmPassword := r.FormValue("confirmPassword")
		if password != ConfirmPassword {
			dialog.Alert("Password and Confirm Password must be same!!")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		} else {

			err := models.Register(username, password)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
