package timeline

type Milestone struct {
	Description string
	ActionType  string
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
