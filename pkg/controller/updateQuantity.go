package controller

import (
	"fmt"
	"lms/pkg/models"
	"net/http"
)

func IncrementBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		quantity := r.FormValue("quantity")
		bookID := r.FormValue("bookID")

		resp := models.IncrementBook(bookID, quantity)
		if resp!=nil {
			fmt.Println((resp))
			http.Error(w, "Failed to process Request", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}

func DecrementBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		quantity := r.FormValue("quantity")
		bookID := r.FormValue("bookID")

		resp := models.DecrementBook(bookID, quantity)
		if resp!=nil {
			fmt.Println((resp))
			http.Redirect(w, r, "/addBook", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}