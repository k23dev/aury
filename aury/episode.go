package aury

import (
	"aury/aury/characters"
	"aury/aury/locations"
	"aury/aury/scenes"
	"fmt"
	"slices"
)

type Episode struct {
	Characters []characters.Character
	Locations  []locations.Location
	Scenes     []scenes.Scene
}

func NewEpisode() *Episode {
	return &Episode{}
}

func (e *Episode) AddScene(name scenes.SceneID) (*scenes.Scene, error) {
	newScene := scenes.New(name)
	// TODO
	e.insertScene(*newScene)
	return newScene, nil
}

func (e *Episode) RemoveScene(id scenes.SceneID) error {
	e.deleteScene(id)
	return nil
}

func (e *Episode) GetScene(id scenes.SceneID) (*scenes.Scene, error) {
	var selectedScene *scenes.Scene
	for _, scene := range e.Scenes {
		if scene.ID == id {
			selectedScene = &scene
			break
		}
	}
	return selectedScene, nil
}

func (e *Episode) AddCharacter(id characters.CharacterID, name string) (*characters.Character, error) {
	newCharacter := characters.NewCharacter(id, name)
	e.insertCharacter(*newCharacter)
	return newCharacter, nil
}

func (e *Episode) RemoveCharacter(id characters.CharacterID) error {
	e.deleteCharacter(id)
	return nil
}

func (e *Episode) GetCharacter(id characters.CharacterID) (*characters.Character, error) {
	var selectedCharacter *characters.Character
	for _, character := range e.Characters {
		if character.ID == id {
			selectedCharacter = &character
			break
		}
	}
	return selectedCharacter, nil
}

func (e *Episode) AddLocation(id locations.LocationID) (*locations.Location, error) {
	newLocation := locations.New()
	e.insertLocation(*newLocation)
	return newLocation, nil
}

func (e *Episode) RemoveLocation(id locations.LocationID) error {
	e.deleteLocation(id)
	return nil
}

func (e *Episode) GetLocation(id locations.LocationID) (locations.Location, error) {
	var selectedLocation locations.Location
	for _, location := range e.Locations {
		if location.ID == id {
			selectedLocation = location
			break
		}
	}
	return selectedLocation, nil
}

// private fn

func (e *Episode) insertScene(newScene scenes.Scene) {
	e.Scenes = append(e.Scenes, newScene)
}

func (e *Episode) deleteScene(oldSceneName scenes.SceneID) {
	auxScenes := slices.Clone(e.Scenes)
	for key, scene := range auxScenes {
		if scene.ID == oldSceneName {
			e.Scenes = slices.Delete(auxScenes, key, 1)
			break
		}
	}
}

func (e *Episode) insertCharacter(newCharacter characters.Character) {
	e.Characters = append(e.Characters, newCharacter)
}

func (e *Episode) deleteCharacter(oldCharacterID characters.CharacterID) {
	auxCharacters := slices.Clone(e.Characters)
	for key, Character := range auxCharacters {
		if Character.ID == oldCharacterID {
			e.Characters = slices.Delete(auxCharacters, key, 1)
			break
		}
	}
}

func (e *Episode) insertLocation(newLocation locations.Location) {
	e.Locations = append(e.Locations, newLocation)
}

func (e *Episode) deleteLocation(oldLocationID locations.LocationID) {
	auxLocations := slices.Clone(e.Locations)
	for key, location := range auxLocations {
		if location.ID == oldLocationID {
			e.Locations = slices.Delete(auxLocations, key, 1)
			break
		}
	}
}

// this shows every milestone in the plot
func (e *Episode) Debug() {
	fmt.Println(" ")
	fmt.Println("########")
	fmt.Println("Debug")
	fmt.Println("########")
	fmt.Println(" ")
	for _, scene := range e.Scenes {
		fmt.Println(" ")
		fmt.Printf("Scene > %s \n", scene.ID)
		fmt.Println("----------------------------")
		for _, milestone := range scene.Timeline.Milestones {
			fmt.Printf("Description: \"%s\" | ActionType: \"%v\" \n", milestone.Description, milestone.ActionType)
		}
	}
}

// this shows every milestone in the plot
func (e *Episode) Play() {
	fmt.Println(" ")
	fmt.Println("########")
	fmt.Println("Playing...")
	fmt.Println("########")
	fmt.Println(" ")
	for _, scene := range e.Scenes {
		fmt.Println(" ")
		fmt.Printf("Scene > %s \n", scene.ID)
		fmt.Println("----------------------------")
		for _, milestone := range scene.Timeline.Milestones {
			fmt.Printf("Description: \"%s\" | ActionType: \"%v\" \n", milestone.Description, milestone.ActionType)
		}
	}
}
