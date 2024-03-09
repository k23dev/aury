package timeline

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
	newMilestone := NewMilestone(description, actionType)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) AddMilestoneAction(description string) {
	actionType := "_action"
	newMilestone := NewMilestone(description, actionType)
	t.Milestones = append(t.Milestones, newMilestone)
}

func (t *Timeline) GetLastMilestone() int {
	return len(t.Milestones) - 1
}
