package views

import (
	"html/template"
)

func LoginPage(file string) *template.Template {
	add := ("templates/" + file + ".html")
	temp := template.Must(template.ParseFiles(add))
	return temp
}
