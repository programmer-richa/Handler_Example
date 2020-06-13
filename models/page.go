package models

import (
	"html/template"
	"log"
	"net/http"
)

var Tpl *template.Template

type Page struct {
	Title             string
	FileName          string
	URL               string
	MainLogoFileName  string
	ShortLogoFileName string
}

func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Here : ", p.Title)
	err := Tpl.ExecuteTemplate(w, p.FileName+".gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
