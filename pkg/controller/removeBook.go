package controller

import (
	"fmt"
	"lms/pkg/models"
	"net/http"
)

func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		name := r.FormValue("bookname")
		fmt.Println(name)
		resp := models.RemBook(name)
		fmt.Println(resp)
		if resp != "" {
			fmt.Println((resp))
			http.Redirect(w, r, "/addBook", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/addBook", http.StatusSeeOther)
		return
	}
}
