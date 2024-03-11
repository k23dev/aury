package episodes

import (
	"aury/aury/aury_error"
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
	SceneChain []int
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
	if selectedScene.ID == "" {
		return selectedScene, aury_error.New("Get Scene", aury_error.SceneNotFound(id), 0)
	}
	return selectedScene, nil
}

func (e *Episode) GetSceneKey(id scenes.SceneID) int {
	for key, scene := range e.Scenes {
		if scene.ID == id {
			return key
		}
	}
	return 0
}

func (e *Episode) JumpToScene(id scenes.SceneID) (*scenes.Scene, error) {
	return e.GetScene(id)
}

func (e *Episode) AddCharacter(character characters.Character) error {
	e.insertCharacter(character)
	return nil
}

func (e *Episode) UpdateCharacter(character characters.Character) error {
	key, err := e.GetCharacterKey(character.ID)
	if err != nil {
		return err
	}
	e.Characters[key] = character
	return nil
}

func (e *Episode) RemoveCharacter(id characters.CharacterID) error {
	e.deleteCharacter(id)
	return nil
}

func (e *Episode) GetCharacter(id characters.CharacterID) (characters.Character, error) {
	var selectedCharacter characters.Character
	for _, character := range e.Characters {
		if character.ID == id {
			return character, nil
		}
	}
	return selectedCharacter, aury_error.New("Get Character", aury_error.CharacterNotFound(id), 0)
}

func (e *Episode) GetCharacterKey(id characters.CharacterID) (int, error) {
	for key, character := range e.Characters {
		if character.ID == id {
			return key, nil
		}
	}
	return 0, aury_error.New("Get Character key", aury_error.CharacterNotFound(id), 0)
}

func (e *Episode) AddLocation(newLocation locations.Location) error {
	e.insertLocation(newLocation)
	return nil
}

func (e *Episode) UpdateLocation(location locations.Location) error {
	key, err := e.GetLocationKey(location.ID)
	if err != nil {
		return err
	}
	e.Locations[key] = location
	return nil
}

func (e *Episode) RemoveLocation(id locations.LocationID) error {
	e.deleteLocation(id)
	return nil
}

func (e *Episode) GetLocation(id locations.LocationID) (locations.Location, error) {
	var selectedLocation locations.Location
	for _, location := range e.Locations {
		if location.ID == id {
			return location, nil
		}
	}
	return selectedLocation, aury_error.New("Get Location", aury_error.LocationNotFound(id), 0)
}

func (e *Episode) GetLocationKey(id locations.LocationID) (int, error) {
	for key, location := range e.Locations {
		if location.ID == id {
			return key, nil
		}
	}
	return 0, aury_error.New("Get Location key", aury_error.LocationNotFound(id), 0)
}

// private fn

func (e *Episode) insertScene(newScene scenes.Scene) {
	e.Scenes = append(e.Scenes, newScene)
}

func (e *Episode) deleteScene(oldSceneID scenes.SceneID) error {
	auxScenes := slices.Clone(e.Scenes)
	for key, scene := range auxScenes {
		if scene.ID == oldSceneID {
			e.Scenes = slices.Delete(auxScenes, key, 1)
			return nil
		}
	}
	return aury_error.New("Deleting scene", aury_error.SceneNotFound(oldSceneID), 0)
}

func (e *Episode) insertCharacter(newCharacter characters.Character) {
	e.Characters = append(e.Characters, newCharacter)
}

func (e *Episode) deleteCharacter(oldCharacterID characters.CharacterID) error {
	auxCharacters := slices.Clone(e.Characters)
	for key, Character := range auxCharacters {
		if Character.ID == oldCharacterID {
			e.Characters = slices.Delete(auxCharacters, key, 1)
			return nil
		}
	}
	return aury_error.New("deleting character", aury_error.CharacterNotFound(oldCharacterID), 0)
}

func (e *Episode) insertLocation(newLocation locations.Location) {
	e.Locations = append(e.Locations, newLocation)
}

func (e *Episode) deleteLocation(oldLocationID locations.LocationID) error {
	auxLocations := slices.Clone(e.Locations)
	for key, location := range auxLocations {
		if location.ID == oldLocationID {
			e.Locations = slices.Delete(auxLocations, key, 1)
			return nil
		}
	}
	return aury_error.New("deleting location", aury_error.LocationNotFound(oldLocationID), 0)
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
		for mKey, milestone := range scene.Timeline.Milestones {
			fmt.Printf("[%d] >> Description: \"%s\" | Type > Code: \"%v\" >  \"%d\" \n", mKey, milestone.Description, milestone.Type, milestone.Code)
		}
	}
}
