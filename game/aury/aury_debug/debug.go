package auryDebug

import (
	"aury/game/aury/characters"
	"aury/game/aury/episodes"
	"aury/game/aury/locations"
	"fmt"
	"strings"
)

func banner() {
	fmt.Println(" ")
	fmt.Println("################")
	fmt.Println("	Debug")
	fmt.Println("################")
	fmt.Println(" ")
}

func separator(text string) {
	banner()
	fmt.Println(" ")
	fmt.Printf("		%s \n", strings.ToUpper(text))
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println(" ")
}

func separator2(text string) {
	fmt.Println(" ")
	fmt.Printf("%s", strings.ToUpper(text))
	separator3()
}

func separator3() {
	fmt.Println("--------------------------------------------------------")
}

// this shows every milestone in the plot
func Episode(e *episodes.Episode) {
	separator("scenes")
	for key, scene := range e.Scenes {
		fmt.Println(" ")
		fmt.Printf(">> [%d] \"%s\" \n", key, scene.ID)
		separator3()
		for mKey, milestone := range scene.Timeline.Milestones {
			fmt.Printf(">>> [%d] >> Description: \"%s\" | Type > Code: \"%v\" >  \"%d\" \n", mKey, milestone.Description, milestone.Type, milestone.Code)
		}
	}
}

// this shows every milestone in the plot
func Character(character *characters.Character) {
	separator("character")
	fmt.Printf(">> ID: \"%s\" Name: \"%s\" \n", character.ID, character.Name)
	// separator2(">>> attributes")
	separator3()
	for sKey, sprite := range character.Sprites {
		fmt.Printf(">>> Sprites > [%d] \"%s\" (%s) = %v \n", sKey, sprite.ID, sprite.ID, sprite.Filepath)
	}
	separator3()
}

// this shows every milestone in the plot
func Characters(characters *[]characters.Character) {
	separator("characters")
	for key, character := range *characters {
		fmt.Println(" ")
		fmt.Printf("[%d] \n", key)
		fmt.Printf(">> ID: \"%s\" Name: \"%s\" \n", character.ID, character.Name)
		fmt.Printf(">> CurrentSprite: %s \n", character.CurrentSprite.ID)
		// separator2(">>> attributes")
		for aKey, attr := range character.Attributes {
			fmt.Printf(">>> Attr > [%d] \"%s\" (%s) = %v \n", aKey, attr.ID, attr.Type, attr.Value)
		}
		separator3()
		for sKey, sprite := range character.Sprites {
			fmt.Printf(">>> Sprites > [%d] \"%s\" (%s) = %v \n", sKey, sprite.ID, sprite.ID, sprite.Filepath)
		}
		separator3()
	}

}

// locations informations

// this shows every milestone in the plot
func Locations(locations *[]locations.Location) {
	separator("locations")
	for key, location := range *locations {
		fmt.Println(" ")
		fmt.Printf("[%d] \n", key)
		fmt.Printf(">> ID: \"%s\" Name: \"%s\" \n", location.ID, location.Name)
		fmt.Printf(">>> %+v \n", location)
		separator3()
	}

}
