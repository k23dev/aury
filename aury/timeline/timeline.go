package timeline

import "aury/aury/characters"

type Timeline struct {
	Milestones []Milestone
}

func New() *Timeline {
	return &Timeline{}
}

func (t *Timeline) AddMilestone(description string) {
	newMilestone := NewMilestone(description, "")
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneDialog(description string) {
	actionType := "_dialog"
	newMilestone := NewMilestone(description, actionType)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneDialogWithOptions(description string) {
	actionType := "_dialog_with_options"
	newMilestone := NewMilestoneAction(description, actionType)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneAction(description string, action interface{}) {
	actionType := "_action"
	newMilestone := NewMilestoneAction(description, actionType)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneActionOfCharacter(description string, character characters.CharacterID, action interface{}) {
	actionType := "_actionOfCharacter"
	newMilestone := NewMilestoneActionFrom(description, actionType, character)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneActionFromCharacterTo(description string, character characters.CharacterID, action interface{}, reciver interface{}) {
	actionType := "_actionFromCharacterTo"
	newMilestone := NewMilestoneActionFromTo(description, actionType, character, reciver)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) GetLastMilestone() int {
	return len(t.Milestones) - 1
}
