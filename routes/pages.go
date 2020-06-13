package routes

import "github.com/programmer-richa/Handler_Example/models"

var LandingPage = models.Page{Title: "Index", FileName: "index", URL: "/"}
var SignupPage = models.Page{Title: "Signup", FileName: "signup", URL: "/signup"}

var Pages = []models.Page{
	LandingPage,
	SignupPage,
}
