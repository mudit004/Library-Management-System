package controller

import (
	"fmt"
	"lms/pkg/models"
	"net/http"
)

func RemBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		name := r.FormValue("bookname")

		resp := models.RemBook(name)

		if resp != "" {
			fmt.Println((resp))
			http.Error(w, "Failed to add book", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
}
