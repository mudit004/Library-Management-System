package controller

import (
	"lms/pkg/models"
	"net/http"
)

func AcceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Extract form values
		UID := r.FormValue("requestUID")
		BID := r.FormValue("requestBID")
		resp := models.AcceptRequest(UID, BID)
		if resp != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/requestManagement", http.StatusSeeOther)
		return
	}

}
