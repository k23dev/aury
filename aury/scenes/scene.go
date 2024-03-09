package scenes

import (
	"aury/aury/characters"
	"aury/aury/dialogs"
	"aury/aury/locations"
	"aury/aury/timeline"
	"fmt"
)

type Scene struct {
	ID       SceneID
	Location *locations.Location
	Music    *locations.Music
	Dialogs  *[]dialogs.Dialog
	Timeline *timeline.Timeline
}

type SceneID string

func New(id SceneID) *Scene {
	timeline := timeline.New()
	music := locations.NewMusic("./default.mp3")
	return &Scene{
		ID:       id,
		Timeline: timeline,
		Music:    music,
	}
}

func (s *Scene) SetMusic(filepath string) error {
	s.Timeline.AddMilestone("new music")

	// TODO
	return nil
}

func (s *Scene) SetLocation(location locations.Location) error {
	description := fmt.Sprintf("location set \"%v\"", location.ID)
	s.Timeline.AddMilestone(description)
	// TODO
	return nil
}

func (s *Scene) AddDialog(character characters.CharacterID, text string) {
	description := fmt.Sprintf("new dialog of character \"%s\"", character)
	s.Timeline.AddMilestoneDialog(description)
	// todo
}

func (s *Scene) AddDialogWithOptions(character characters.CharacterID, text string) {
	description := fmt.Sprintf("new dialog with options of character \"%s\"", character)
	s.Timeline.AddMilestoneDialogWithOptions(description, nil)
	// todo
}

func (s *Scene) AddActionOfCharacter(description string, action interface{}, character characters.CharacterID) {
	description = fmt.Sprintf("new action: \"%v\" from \"%s\"", description, character)
	s.Timeline.AddMilestoneActionOfCharacter(description, action, character)
	// todo
}

func (s *Scene) AddActionFromCharacterTo(description string, action interface{}, character characters.CharacterID, reciver interface{}) {
	description = fmt.Sprintf("new action: \"%v\" from \"%s\" -> \"%s\"", description, character, reciver)
	s.Timeline.AddMilestoneActionFromCharacterTo(description, action, character, reciver)
	// todo
}

func (s *Scene) AddAction(description string, action interface{}) {
	description = fmt.Sprintf("new action: \"%v\"", action)
	s.Timeline.AddMilestoneAction(description, action)
	// todo
}

func (s *Scene) PlayScene() {
	for key, milestone := range s.Timeline.Milestones {
		fmt.Printf("%d > %s ", key, milestone.Description)
	}
}

// this function must wait until some action happens
func (s *Scene) AwaitScene() {
	// TODO:
}

func (s *Scene) PlayMusic() {
	s.Timeline.AddMilestone("Playing music")
	// todo
}
