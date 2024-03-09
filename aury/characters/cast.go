package characters

type Cast struct {
	Characters *[]Character
}

func NewCast() *Cast {
	return &Cast{}
}

func (c *Cast) AddCharacter() error {
	// TODO
	return nil
}

func (c *Cast) RemoveCharacter() error {
	// TODO
	return nil
}

func (c *Cast) GetCharacter(ID string) error {
	// TODO
	return nil
}
