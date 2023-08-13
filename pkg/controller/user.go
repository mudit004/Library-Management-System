package controller

import (
	"fmt"
	"lms/pkg/models"
	"lms/pkg/utils"
	"lms/pkg/views"
	"net/http"
)

func Browse(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("Browse")
	writer.WriteHeader(http.StatusOK)

	formData := models.FetchBook(utils.GetUID(writer, request))
	t.Execute(writer, formData)
}

func RequestBookLoader(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		Logout(writer, request)
		return
	} else {
		t := views.LoginPage("requestBook")
		writer.WriteHeader(http.StatusOK)

		formData := models.FetchBook(utils.GetUID(writer, request))
		t.Execute(writer, formData)

	}

}

func ReturnBookLoader(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		Logout(writer, request)
		return
	} else {
		t := views.LoginPage("returnBook")
		writer.WriteHeader(http.StatusOK)
		fmt.Println("#####################################", utils.GetUID(writer, request))
		formData := models.FetchBook(utils.GetUID(writer, request))
		t.Execute(writer, formData)

	}

}

func PendingRequestLoader(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		Logout(writer, request)
		return
	} else {
		t := views.LoginPage("requestPending")
		writer.WriteHeader(http.StatusOK)

		formData := models.FetchBook(utils.GetUID(writer, request))
		t.Execute(writer, formData)

	}

}
