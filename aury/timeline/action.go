package timeline

import (
	"aury/aury/characters"
	"aury/aury/dialogs"
)

type Action struct {
	Character *characters.Character
	Dialog    *dialogs.Dialog
}
