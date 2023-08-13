package controller

import (
	"fmt"
	"lms/pkg/models"
	"lms/pkg/utils"
	"lms/pkg/views"
	"net/http"
)

func IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("access-token")
	fmt.Println(cookie)
	if err != nil {
		fmt.Println("Cookie Does Not exist")
		return false
	} else {
		fmt.Println("Hey I am Here")
		token := cookie.Value
		fmt.Println(token)
		claims, err := utils.DecodeJWT(token)
		fmt.Println(claims)
		if err != nil {
			fmt.Println("Error in decoding")
			return false
		} else {
			fmt.Println(claims["admin"] == float64(1))
			if claims["admin"] == float64(1) {
				fmt.Println("returning true")
				return true
			} else {
				fmt.Println("returning false")
				return false
			}
		}
	}
}

func AddBookPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Mere Pyaar Bhaiya aur Beheno ", IsAdmin(writer, request))
	if IsAdmin(writer, request) {
		fmt.Println("Started AddBookPage")
		t := views.LoginPage("Books")
		fmt.Println(t)
		writer.WriteHeader(http.StatusOK)
		formData := models.FetchBook(utils.GetUID(writer, request))
		fmt.Println(formData)
		t.Execute(writer, formData)
	} else {
		Logout(writer, request)

	}
}

func RequestManagementPage(writer http.ResponseWriter, request *http.Request) {
	if IsAdmin(writer, request) {
		fmt.Println("in if")
		t := views.LoginPage("request")
		writer.WriteHeader(http.StatusOK)
		formData := models.FetchBook(utils.GetUID(writer, request))
		t.Execute(writer, formData)
	} else {
		fmt.Println("in else")
		Logout(writer, request)
	}
}
