package controller

import (
	"lms/pkg/models"
	"net/http"
)

func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		name := r.FormValue("bookname")
		resp := models.RemoveBook(name)
		if resp != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}
