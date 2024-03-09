package main

import (
	"aury/myplot/episode1"
)

func main() {
	episode1.Setup()
	ep1 := episode1.Episode()
	ep1.Debug()

}
