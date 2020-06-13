package main

import (
	"html/template"
	"net/http"

	"github.com/programmer-richa/Handler_Example/models"
	"github.com/programmer-richa/Handler_Example/routes"
)

// Initializer
func init() {
	models.Tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}
func main() {
	for _, page := range routes.Pages {
		http.Handle(page.URL, page)
	}
	// http.Handle("/", models.LandingPage)
	// http.Handle("/signup", models.SignupPage)
	// http.ListenAndServe("localhost:8080/signup", models.SignupPage)
	// Listener
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServe(":8080", models.LandingPage)
}
