package aury

import (
	"aury/aury/characters"
	"aury/aury/locations"
	"aury/aury/scenes"
	"slices"
)

type Plot struct {
	Characters *[]characters.Character
	Locations  *[]locations.Location
	Scenes     *[]scenes.Scene
}

func NewPlot() *Plot {
	return &Plot{}
}

func (p *Plot) AddScene(name scenes.SceneID) (*scenes.Scene, error) {
	newScene := scenes.New(name)
	// TODO
	p.insertScene(*newScene)
	return newScene, nil
}

func (p *Plot) RemoveScene(id scenes.SceneID) error {
	p.deleteScene(id)
	return nil
}

func (p *Plot) GetScene(id scenes.SceneID) (*scenes.Scene, error) {
	var selectedScene *scenes.Scene
	for _, scene := range *p.Scenes {
		if scene.ID == id {
			selectedScene = &scene
			break
		}
	}
	return selectedScene, nil
}

func (p *Plot) AddCharacter(id characters.CharacterID, name string) (*characters.Character, error) {
	newCharacter := characters.NewCharacter(id, name)
	p.insertCharacter(*newCharacter)
	return newCharacter, nil
}

func (p *Plot) RemoveCharacter(id characters.CharacterID) error {
	p.deleteCharacter(id)
	return nil
}

func (p *Plot) GetCharacter(id characters.CharacterID) (*characters.Character, error) {
	var selectedCharacter *characters.Character
	for _, character := range *p.Characters {
		if character.ID == id {
			selectedCharacter = &character
			break
		}
	}
	return selectedCharacter, nil
}

func (p *Plot) AddLocation(id locations.LocationID) (*locations.Location, error) {
	newLocation := locations.New()
	p.insertLocation(*newLocation)
	return newLocation, nil
}

func (p *Plot) RemoveLocation(id locations.LocationID) error {
	p.deleteLocation(id)
	return nil
}

func (p *Plot) GetLocation(id locations.LocationID) (*locations.Location, error) {
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

func (p *Plot) insertScene(newScene scenes.Scene) {
	scenes := append(*p.Scenes, newScene)
	p.Scenes = &scenes
}

func (p *Plot) deleteScene(oldSceneName scenes.SceneID) {
	auxScenes := slices.Clone(*p.Scenes)
	for key, scene := range auxScenes {
		if scene.ID == oldSceneName {
			auxScenes = slices.Delete(auxScenes, key, 1)
			p.Scenes = &auxScenes
			break
		}
	}
}

func (p *Plot) insertCharacter(newCharacter characters.Character) {
	characters := append(*p.Characters, newCharacter)
	p.Characters = &characters
}

func (p *Plot) deleteCharacter(oldCharacterID characters.CharacterID) {
	auxCharacters := slices.Clone(*p.Characters)
	for key, Character := range auxCharacters {
		if Character.ID == oldCharacterID {
			auxCharacters = slices.Delete(auxCharacters, key, 1)
			p.Characters = &auxCharacters
			break
		}
	}
}

func (p *Plot) insertLocation(newLocation locations.Location) {
	locations := append(*p.Locations, newLocation)
	p.Locations = &locations
}

func (p *Plot) deleteLocation(oldLocationID locations.LocationID) {
	auxLocations := slices.Clone(*p.Locations)
	for key, location := range auxLocations {
		if location.ID == oldLocationID {
			auxLocations = slices.Delete(auxLocations, key, 1)
			p.Locations = &auxLocations
			break
		}
	}
}
