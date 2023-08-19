package controller

import (
	"lms/pkg/models"
	"lms/pkg/utils"
	"lms/pkg/views"
	"net/http"
)

func Browse(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage("Browse")
	writer.WriteHeader(http.StatusOK)

	userID, err := utils.GetUserID(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
		return

	}
	formData, err := models.FetchBook(userID)
	if err != nil {
		http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
		return

	}
	t.Execute(writer, formData)
}

func RequestBookLoader(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		Logout(writer, request)
		return
	} else {
		t := views.LoginPage("requestBook")
		writer.WriteHeader(http.StatusOK)
		userID, err := utils.GetUserID(writer, request)
		if err != nil {
			http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
			return

		}
		formData, err := models.FetchBook(userID)
		if err != nil {
			http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
			return

		}
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
		userID, err := utils.GetUserID(writer, request)
		if err != nil {
			http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
			return

		}
		formData, err := models.FetchBook(userID)
		if err != nil {
			http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
			return

		}
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
		userID, err := utils.GetUserID(writer, request)
		if err != nil {
			http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
			return

		}
		formData, err := models.FetchBook(userID)
		if err != nil {
			http.Redirect(writer, request, "/internalServerError", http.StatusSeeOther)
			return

		}
		t.Execute(writer, formData)

	}

}
