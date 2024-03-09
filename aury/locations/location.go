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

func (l *Location) SetName(name string) error {
	l.Name = name
	return nil
}

func (l *Location) SetBackground(filepath string) error {
	l.Background = filepath
	return nil
}

func (l *Location) AddBgStyle(style string) error {
	l.insertStyle(style)
	return nil
}

func (l *Location) ResetBgStyle() error {
	var reset []string
	l.BgStyle = reset
	return nil
}

func (l *Location) SetBluriedBgStyle() {
	// TODO
}

func (l *Location) SetShapedBgStyle() {
	// TODO
}

func (l *Location) insertStyle(style string) {
	l.BgStyle = append(l.BgStyle, style)
}
