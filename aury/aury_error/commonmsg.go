package auryError

import (
	"aury/aury/characters"
	"aury/aury/locations"
	"aury/aury/scenes"
	"fmt"
)

func SceneNotFound(id scenes.SceneID) string {
	return fmt.Sprintf("Scene %s not found", id)
}

func CharacterNotFound(id characters.CharacterID) string {
	return fmt.Sprintf("Character %s not found", id)
}

func LocationNotFound(id locations.LocationID) string {
	return fmt.Sprintf("Location %s not found", id)
}
