package dialogs

import "aury/aury/characters"

type Dialog struct {
	CharacterID characters.CharacterID
	Text        string
}

func NewDialog(characterID characters.CharacterID, text string) *Dialog {
	return &Dialog{
		CharacterID: characterID,
		Text:        text,
	}
}
