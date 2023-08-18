package controller

import (
	"lms/pkg/models"
	"lms/pkg/utils"
	"lms/pkg/views"
	"net/http"
)

func IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		return false
	} else {
		token := cookie.Value
		claims, err := utils.DecodeJWT(token)
		if err != nil {
			return false
		} else {
			if claims["admin"] == float64(1) {
				return true
			} else {
				return false
			}
		}
	}
}

func AddBookPage(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		t := views.LoginPage("Books")
		writer.WriteHeader(http.StatusOK)
		formData := models.FetchBook(utils.GetUID(writer, request))
		t.Execute(writer, formData)
	} else {
		Logout(writer, request)

	}
}

func RequestManagementPage(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		t := views.LoginPage("request")
		writer.WriteHeader(http.StatusOK)
		formData := models.FetchBook(utils.GetUID(writer, request))
		t.Execute(writer, formData)
	} else {
		Logout(writer, request)
	}
}
