package characters

// TODO
type Character struct {
	ID         CharacterID
	Name       string
	Attributes []Attribute
}

type CharacterID string

func NewCharacter(id CharacterID, name string) *Character {
	return &Character{
		ID:   id,
		Name: name,
	}
}

func (c *Character) AddAttribute() error {
	// TODO
	return nil
}

func (c *Character) RemoveAttribute() error {
	// TODO
	return nil
}
