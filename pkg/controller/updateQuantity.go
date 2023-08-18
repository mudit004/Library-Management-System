package controller

import (
	"lms/pkg/models"
	"net/http"
)

func IncrementBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		quantity := r.FormValue("quantity")
		bookID := r.FormValue("bookID")

		resp := models.IncrementBook(bookID, quantity)
		if resp != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}

func DecrementBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		quantity := r.FormValue("quantity")
		bookID := r.FormValue("bookID")

		_ = models.DecrementBook(bookID, quantity)

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}
