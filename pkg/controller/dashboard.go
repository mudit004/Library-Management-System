package controller

import (
	"lms/pkg/models"
	"lms/pkg/views"
	"net/http"
)

func Dashboard(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("dashboard")
	writer.WriteHeader(http.StatusOK)

	formData := models.FetchBook()
	t.Execute(writer, formData)
}
