package controller

import (
	"lms/pkg/views"
	"net/http"
)

func ServerUnavailable(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	t := views.RenderPage("503error")
	t.Execute(w, nil)

}
