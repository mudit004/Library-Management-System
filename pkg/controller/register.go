package controller

import (
	"fmt"
	"lms/pkg/models"
	"lms/pkg/views"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("Register")
	t.Execute(writer, nil)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		Cpassword := r.FormValue("confirmPassword")
		fmt.Println(username)
		if password != Cpassword {
			fmt.Println("Confirm Pwd does not match")
			http.Redirect(w, r, "/register", http.StatusSeeOther)

		}

		err := models.Register(username, password)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/register", http.StatusSeeOther)

		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
