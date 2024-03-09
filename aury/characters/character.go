package characters

import (
	"fmt"
	"slices"
)

// TODO
type Character struct {
	ID            CharacterID
	Name          string
	Attributes    []Attribute
	Sprites       []Sprite
	CurrentSprite Sprite
}

type CharacterID string

func NewCharacter(id CharacterID, name string) *Character {
	return &Character{
		ID:   id,
		Name: name,
	}
}

// AddAttribute adds a new attribute to the provided slice.
func (c *Character) AddAttribute(id AttributeID, typ string, value interface{}) {
	newAttribute := Attribute{
		ID:    id,
		Type:  typ,
		Value: value,
	}
	c.insertAttribute(newAttribute)
}

// EditAttribute edits the value of an attribute in the provided slice.
func (c *Character) EditAttribute(id AttributeID, newValue interface{}) {
	for i := range c.Attributes {
		if c.Attributes[i].ID == id {
			c.Attributes[i].Value = newValue
			break
		}
	}
}

// ReadAttribute retrieves and returns the value of an attribute from the provided slice.
func (c *Character) ReadAttribute(id AttributeID) Attribute {
	for _, attr := range c.Attributes {
		if attr.ID == id {
			return attr
		}
	}
	return Attribute{} // Attribute not found
}

func (c *Character) insertAttribute(newAttribute Attribute) {
	c.Attributes = append(c.Attributes, newAttribute)
}

func (c *Character) DeleteAttribute(id AttributeID) error {
	for key, attr := range c.Attributes {
		if attr.ID == id {
			c.Attributes = slices.Delete(c.Attributes, key, 1)
			return nil
		}
	}

	return fmt.Errorf("attribute not found")
}

func (c *Character) AddSprite(id SpriteID, filepath string) error {
	newSprite := NewSprite(id, filepath)
	c.Sprites = append(c.Sprites, *newSprite)
	return nil
}

func (c *Character) SetSprite(id SpriteID) error {
	currentSprite, err := c.GetSprite(id)
	if err != nil {
		return err
	}
	c.CurrentSprite = *currentSprite
	return nil
}

func (c *Character) GetSprite(id SpriteID) (*Sprite, error) {
	selectedSprite := &Sprite{}
	for _, sprite := range c.Sprites {
		if sprite.ID == id {
			return &sprite, nil
		}
	}
	return selectedSprite, fmt.Errorf("Sprite not found")
}

func (c *Character) GetSpriteKey(id SpriteID) (int, error) {
	for key, sprite := range c.Sprites {
		if sprite.ID == id {
			return key, nil
		}
	}
	return 0, fmt.Errorf("Sprite not found")
}
