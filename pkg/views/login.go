package views

import (
	"html/template"
)

func RenderPage(file string) *template.Template {
	add := ("templates/" + file + ".html")
	temp := template.Must(template.ParseFiles(add))
	return temp
}
