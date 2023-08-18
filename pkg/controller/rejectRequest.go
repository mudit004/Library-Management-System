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
		if resp != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/requestManagement", http.StatusSeeOther)
		return
	}

}
