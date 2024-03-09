package main

import (
	auryDebug "aury/aury/aury_debug"
	"aury/myplot/episode1"
)

func main() {
	ep1 := episode1.Episode()
	auryDebug.Episode(&ep1)
	// auryDebug.Character(&ep1.Characters[0])
	auryDebug.Characters(&ep1.Characters)
}
