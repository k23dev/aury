package locations

// TODO
type Location struct {
	ID         LocationID
	Name       string
	Background string
	BgStyle    []string
}

type LocationID string

func New() *Location {
	return &Location{}
}

func (l *Location) SetBackground(filepath string) error {
	// TODO
	return nil
}

func (l *Location) SetBgStyle(filepath string) error {
	// TODO
	return nil
}

func (l *Location) SetBluriedBgStyle() {
	// TODO
}

func (l *Location) SetShapedBgStyle() {
	// TODO
}
