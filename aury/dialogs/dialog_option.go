package dialogs

// This kind is a variant of Dialog.
// the player can chose one option of this dialogo
// every option has one consecuence.
// The consecuence can be scene
type DialogOption struct {
	Text        string
	Consecuence string
}

func NewDialogOption(text string, consecuence string) *DialogOption {
	return &DialogOption{
		Text:        text,
		Consecuence: consecuence,
	}
}
