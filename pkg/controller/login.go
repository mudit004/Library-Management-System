package controller

import (
	"fmt"
	"lms/pkg/models"
	"lms/pkg/views"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("loginPage")
	t.Execute(writer, nil)
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		username := r.FormValue("username")
		password := r.FormValue("pwd")
		fmt.Println(username, " & ", password)

		resp, token, isAdmin := models.Login(username, password)
		if resp != true {
			fmt.Println((resp))
			fmt.Println("Failed to Login")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			Cookie := http.Cookie{
				Name:    "access-token",
				Value:   token,
				Expires: time.Now().Add(48 * time.Hour),
				Path:    "/",
			}
			http.SetCookie(w, &Cookie)
			if isAdmin {
				http.Redirect(w, r, "/addBook", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/requestBook", http.StatusSeeOther)
			}

		}

	}
}
