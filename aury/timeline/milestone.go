package timeline

import "aury/aury/characters"

type Milestone struct {
	Code        uint
	Description string
	Type        string
	Action      *Action
}

func NewMilestone(code uint, description, mType string) Milestone {

	if mType == "" {
		mType = "_"
	}

	return Milestone{
		Description: description,
		Type:        mType,
		Code:        code,
	}
}

func NewMilestoneAction(code uint, description, mType string, action interface{}) Milestone {

	newAction := NewAction(action)

	return Milestone{
		Description: description,
		Type:        mType,
		Code:        code,
		Action:      newAction,
	}
}

// set action maded by a character
func NewMilestoneActionFrom(code uint, description, mType string, action interface{}, characterFrom characters.CharacterID) Milestone {

	newAction := NewActionFrom(action, characterFrom)

	return Milestone{
		Description: description,
		Type:        mType,
		Code:        code,
		Action:      newAction,
	}
}

// set action maded by a character to [character,other]
func NewMilestoneActionFromTo(code uint, description, mType string, action interface{}, characterFrom characters.CharacterID, reciver interface{}) Milestone {

	newAction := NewActionFromTo(action, characterFrom, reciver)

	return Milestone{
		Description: description,
		Type:        mType,
		Code:        code,
		Action:      newAction,
	}
}
