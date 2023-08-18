package controller

import (
	"lms/pkg/models"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		name := r.FormValue("bookname")
		quantity := r.FormValue("quantity")
		resp := models.AddBook(name, quantity)
		if resp != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}
