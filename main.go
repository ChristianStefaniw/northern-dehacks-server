package main

import (
	"github.com/ChristianStefaniw/cgr"

	"github.com/northern-dehacks/northern-dehacks-server/controllers"
	"github.com/northern-dehacks/northern-dehacks-server/database"
	"github.com/northern-dehacks/northern-dehacks-server/helpers"
	"github.com/northern-dehacks/northern-dehacks-server/keys"
	"github.com/northern-dehacks/northern-dehacks-server/middleware"
)

func main() {
	keys.Load()
	database.Connect()
	helpers.AuthEmail()

	port := keys.PORT

	router := cgr.NewRouter()
	corsMiddleware := cgr.NewMiddleware(middleware.CorsMiddleware)

	router.Route("/").Handler(controllers.Home).Method("GET").Insert()
	router.Route("/newsletter").Handler(controllers.SignUpForNewsletter).Method("POST", "OPTIONS").HandlePreflight(true).Assign(corsMiddleware).Insert()
	router.Route("/contact").Handler(controllers.Contact).Method("POST", "OPTIONS").HandlePreflight(true).Assign(corsMiddleware).Insert()

	router.Run(port)
}
