package controller

import (
	"lms/pkg/views"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	t := views.RenderPage("404error")
	t.Execute(w, nil)

}
