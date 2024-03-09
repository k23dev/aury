package aury

import (
	"aury/aury/characters"
	"aury/aury/locations"
	"aury/aury/scenes"
	"fmt"
	"slices"
)

type Episode struct {
	Characters *[]characters.Character
	Locations  *[]locations.Location
	Scenes     *[]scenes.Scene
}

func NewEpisode() *Episode {
	return &Episode{}
}

func (p *Episode) AddScene(name scenes.SceneID) (*scenes.Scene, error) {
	newScene := scenes.New(name)
	// TODO
	p.insertScene(*newScene)
	return newScene, nil
}

func (p *Episode) RemoveScene(id scenes.SceneID) error {
	p.deleteScene(id)
	return nil
}

func (p *Episode) GetScene(id scenes.SceneID) (*scenes.Scene, error) {
	var selectedScene *scenes.Scene
	for _, scene := range *p.Scenes {
		if scene.ID == id {
			selectedScene = &scene
			break
		}
	}
	return selectedScene, nil
}

func (p *Episode) AddCharacter(id characters.CharacterID, name string) (*characters.Character, error) {
	newCharacter := characters.NewCharacter(id, name)
	p.insertCharacter(*newCharacter)
	return newCharacter, nil
}

func (p *Episode) RemoveCharacter(id characters.CharacterID) error {
	p.deleteCharacter(id)
	return nil
}

func (p *Episode) GetCharacter(id characters.CharacterID) (*characters.Character, error) {
	var selectedCharacter *characters.Character
	for _, character := range *p.Characters {
		if character.ID == id {
			selectedCharacter = &character
			break
		}
	}
	return selectedCharacter, nil
}

func (p *Episode) AddLocation(id locations.LocationID) (*locations.Location, error) {
	newLocation := locations.New()
	p.insertLocation(*newLocation)
	return newLocation, nil
}

func (p *Episode) RemoveLocation(id locations.LocationID) error {
	p.deleteLocation(id)
	return nil
}

func (p *Episode) GetLocation(id locations.LocationID) (*locations.Location, error) {
	var selectedLocation *locations.Location
	for _, location := range *p.Locations {
		if location.ID == id {
			selectedLocation = &location
			break
		}
	}
	return selectedLocation, nil
}

// private fn

func (p *Episode) insertScene(newScene scenes.Scene) {
	scenes := append(*p.Scenes, newScene)
	p.Scenes = &scenes
}

func (p *Episode) deleteScene(oldSceneName scenes.SceneID) {
	auxScenes := slices.Clone(*p.Scenes)
	for key, scene := range auxScenes {
		if scene.ID == oldSceneName {
			auxScenes = slices.Delete(auxScenes, key, 1)
			p.Scenes = &auxScenes
			break
		}
	}
}

func (p *Episode) insertCharacter(newCharacter characters.Character) {
	characters := append(*p.Characters, newCharacter)
	p.Characters = &characters
}

func (p *Episode) deleteCharacter(oldCharacterID characters.CharacterID) {
	auxCharacters := slices.Clone(*p.Characters)
	for key, Character := range auxCharacters {
		if Character.ID == oldCharacterID {
			auxCharacters = slices.Delete(auxCharacters, key, 1)
			p.Characters = &auxCharacters
			break
		}
	}
}

func (p *Episode) insertLocation(newLocation locations.Location) {
	locations := append(*p.Locations, newLocation)
	p.Locations = &locations
}

func (p *Episode) deleteLocation(oldLocationID locations.LocationID) {
	auxLocations := slices.Clone(*p.Locations)
	for key, location := range auxLocations {
		if location.ID == oldLocationID {
			auxLocations = slices.Delete(auxLocations, key, 1)
			p.Locations = &auxLocations
			break
		}
	}
}

// this shows every milestone in the plot
func (p *Episode) Debug() {
	for _, scene := range *p.Scenes {
		for _, milestone := range scene.Timeline.Milestones {
			fmt.Printf("%v \n", milestone)
		}
	}
}
