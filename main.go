package main

import (
	"aury/screens"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	// ep1 := episode1.Episode()
	// // ep2 := episode2.Episode()

	// // debug episode
	// auryDebug.Locations(&ep1.Locations)
	// auryDebug.Episode(&ep1)
	// // auryDebug.Episode(&ep2)

	// // debug characters
	// // auryDebug.Character(&ep1.Characters[0])
	// auryDebug.Characters(&ep1.Characters)
	// // auryDebug.Characters(&ep2.Characters)

	app.Route("/", &screens.Hello{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
