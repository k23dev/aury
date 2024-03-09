package scenes

import (
	"aury/aury/characters"
	"aury/aury/dialogs"
	"aury/aury/locations"
	"aury/aury/timeline"
	"fmt"
)

type Scene struct {
	ID SceneID
	// Prev       *Scene
	// Next       *[]Scene
	Location *locations.Location
	Music    *locations.Music
	// Characters *[]characters.Character
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

func (s *Scene) SetLocation(location *locations.Location) error {
	description := fmt.Sprintf("location set %s\n", location.ID)
	s.Timeline.AddMilestone(description)
	// TODO
	return nil
}

func (s *Scene) AddDialog(character *characters.CharacterID, text string) {
	description := fmt.Sprintf("new dialog of character %s\n", *character)
	s.Timeline.AddMilestoneDialog(description)
	// todo
}

func (s *Scene) AddDialogWithOptions(character *characters.CharacterID, text string) {
	description := fmt.Sprintf("new dialog with options of character %s\n", *character)
	s.Timeline.AddMilestoneDialogWithOptions(description)
	// todo
}

// func (s *Scene) SetPrevScene(prev *Scene) error {
// 	s.Prev = prev
// 	// TODO
// 	return nil
// }

// func (s *Scene) AddNextScene(next *Scene) error {
// 	// TODO
// 	return nil
// }

// func (s *Scene) RemoveNextScene(next *Scene) error {
// 	// TODO
// 	return nil
// }

func (s *Scene) PlayScene() {
	for key, milestone := range s.Timeline.Milestones {
		fmt.Printf("%d > %s \n", key, milestone.Description)
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
