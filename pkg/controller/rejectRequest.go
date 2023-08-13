package controller

import (
	"fmt"
	"lms/pkg/models"
	"net/http"
)

func RejectRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		UID := r.FormValue("requestUID")
		BID := r.FormValue("requestBID")
		fmt.Println(UID, " & ", BID)

		resp := models.RejectRequest(UID, BID)
		if resp != "" {
			fmt.Println((resp))
			http.Error(w, "Failed to reject request", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

}
