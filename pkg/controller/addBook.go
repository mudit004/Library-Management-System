package controller

import (
	"fmt"
	"lms/pkg/models"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		name := r.FormValue("bookname")
		quantity := r.FormValue("quantity")
		fmt.Println(name, " & ", quantity)

		resp := models.AddBook(name, quantity)
		if resp != "" {
			fmt.Println((resp))
			http.Error(w, "Failed to add book", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
}
