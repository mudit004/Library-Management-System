package controller

import (
	"lms/pkg/models"
	"lms/pkg/views"
	"net/http"
)

func Browse(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("Browse")
	writer.WriteHeader(http.StatusOK)

	formData := models.FetchBook()
	t.Execute(writer, formData)
}
