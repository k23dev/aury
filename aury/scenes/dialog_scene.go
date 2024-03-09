package scenes

import (
	"aury/aury/characters"
	"aury/aury/dialogs"
)

type SceneDialog struct {
	Character *characters.Character
	Dialog    *dialogs.Dialog
}
