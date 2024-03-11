package timeline

import (
	"aury/game/aury/characters"
	"aury/game/aury/milestonescodes"
)

type Timeline struct {
	Milestones []Milestone
}

func New() *Timeline {
	return &Timeline{}
}

func (t *Timeline) AddMilestone(description string) {
	newMilestone := NewMilestone(milestonescodes.Generic, description, "")
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneLocation(description string) {
	newMilestone := NewMilestone(milestonescodes.LocationGeneric, description, "")
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneDialog(description string, characterID characters.CharacterID, text string) {
	actionType := "_dialog"
	newMilestone := NewMilestoneDialog(milestonescodes.DialogGeneric, description, actionType, characterID, text)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneDialogWithOptions(description string, action interface{}) {
	actionType := "_dialog_with_options"
	newMilestone := NewMilestoneAction(milestonescodes.DialogWithOptions, description, actionType, action)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneAction(description string, action interface{}) {
	actionType := "_action"
	newMilestone := NewMilestoneAction(milestonescodes.ActionGeneric, description, actionType, action)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneActionOfCharacter(description string, action interface{}, character characters.CharacterID) {
	actionType := "_actionOfCharacter"
	newMilestone := NewMilestoneActionFrom(milestonescodes.ActionOf, description, actionType, action, character)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneActionFromCharacterTo(description string, action interface{}, character characters.CharacterID, reciver interface{}) {
	actionType := "_actionFromCharacterTo"
	newMilestone := NewMilestoneActionFromTo(milestonescodes.ActionFromTo, description, actionType, action, character, reciver)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) GetLastMilestone() int {
	return len(t.Milestones) - 1
}

func (t *Timeline) CountMilestones() int {
	return len(t.Milestones)
}
