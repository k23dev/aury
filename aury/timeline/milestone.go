package timeline

import "aury/aury/characters"

type Milestone struct {
	Description string
	ActionType  string
	Action      interface{}
	ActionFrom  characters.CharacterID
	ActionTo    interface{}
}

func NewMilestone(description, actionType string) Milestone {

	if actionType == "" {
		actionType = "null"
	}

	return Milestone{
		Description: description,
		ActionType:  actionType,
	}
}

func NewMilestoneAction(description, actionType string) Milestone {

	return Milestone{
		Description: description,
		ActionType:  actionType,
	}
}

// set action maded by a character
func NewMilestoneActionFrom(description, actionType string, characterFrom characters.CharacterID) Milestone {

	return Milestone{
		Description: description,
		ActionType:  actionType,
		ActionFrom:  characterFrom,
	}
}

// set action maded by a character to [character,other]
func NewMilestoneActionFromTo(description, actionType string, characterFrom characters.CharacterID, reciver interface{}) Milestone {

	return Milestone{
		Description: description,
		ActionType:  actionType,
		ActionFrom:  characterFrom,
		ActionTo:    reciver,
	}
}

func (m *Milestone) SetActionFrom(character characters.CharacterID) {
	m.ActionFrom = character
}

func (m *Milestone) SetActionTo(reciver interface{}) {
	m.ActionTo = reciver
}

func (m *Milestone) SetAction(action interface{}) {
	m.Action = action
}
