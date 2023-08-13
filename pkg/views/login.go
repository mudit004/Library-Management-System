package views

import (
	"fmt"
	"html/template"
)

func LoginPage(file string) *template.Template {
	add := ("templates/" + file + ".html")
	fmt.Println(add)
	temp := template.Must(template.ParseFiles(add))
	return temp
}
