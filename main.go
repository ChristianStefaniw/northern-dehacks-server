package main

import (
	"github.com/ChristianStefaniw/cgr"
	"github.com/northern-dehacks/northern-dehacks-server/controllers"
	"github.com/northern-dehacks/northern-dehacks-server/database"
	"github.com/northern-dehacks/northern-dehacks-server/keys"
)

func main() {
	keys.Load()
	database.Connect()

	port := keys.PORT

	router := cgr.NewRouter()
	router.Route("/").Handler(controllers.Home).Method("GET").Insert()
	router.Route("/newsletter").Handler(controllers.SignUpForNewsletter).Method("POST", "OPTIONS").HandlePreflight(true).Insert()

	router.Run(port)
}
