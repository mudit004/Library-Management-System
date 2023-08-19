package controller

import (
	"lms/pkg/models"
	"lms/pkg/views"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	t := views.RenderPage("loginPage")
	t.Execute(writer, nil)
}

func Admin(writer http.ResponseWriter, request *http.Request) {
	t := views.RenderPage("loginAdmin")
	t.Execute(writer, nil)
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		resp, token, isAdmin, err := models.Login(username, password)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}
		if !resp {
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
				cookie := http.Cookie{
					Name:    "access-token",
					Value:   "",
					Expires: time.Now(),
				}

				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/admin", http.StatusSeeOther)

			} else {
				http.Redirect(w, r, "/requestBook", http.StatusSeeOther)
			}

		}

	}
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		resp, token, isAdmin, err := models.Login(username, password)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}

		if !resp {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			cookie := http.Cookie{
				Name:    "access-token",
				Value:   token,
				Expires: time.Now().Add(48 * time.Hour),
				Path:    "/",
			}
			http.SetCookie(w, &cookie)
			if isAdmin {
				http.Redirect(w, r, "/addBook", http.StatusSeeOther)
			} else {
				cookie := http.Cookie{
					Name:    "access-token",
					Value:   "",
					Expires: time.Now(),
				}

				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/", http.StatusSeeOther)

			}

		}

	}
}
