package timeline

import "aury/aury/characters"

type Action struct {
	Action interface{}
	From   characters.CharacterID
	To     interface{}
}

func NewAction(action interface{}) *Action {
	return &Action{
		Action: action,
	}
}

func NewActionFrom(action interface{}, from characters.CharacterID) *Action {
	return &Action{
		Action: action,
		From:   from,
	}
}

func NewActionFromTo(action interface{}, from characters.CharacterID, to interface{}) *Action {
	return &Action{
		Action: action,
		From:   from,
		To:     to,
	}
}
