package main

import (
	"aury/myplot/episode1"
	"fmt"
)

func main() {
	episode1.Setup()
	ep1 := episode1.Episode()
	for _, scene := range *ep1 {
		for _, milestone := range scene.Timeline.Milestones {
			fmt.Printf("%v \n", milestone)
		}
	}
}
