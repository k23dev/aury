package main

import (
	auryDebug "aury/aury/aury_debug"
	"aury/myplot/episode1"
)

func main() {
	ep1 := episode1.Episode()
	// ep2 := episode2.Episode()

	// debug episode
	auryDebug.Episode(&ep1)
	auryDebug.Locations(&ep1.Locations)
	// auryDebug.Episode(&ep2)

	// debug characters
	// auryDebug.Character(&ep1.Characters[0])
	auryDebug.Characters(&ep1.Characters)
	// auryDebug.Characters(&ep2.Characters)
}
