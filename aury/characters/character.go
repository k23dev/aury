package characters

import (
	"fmt"
	"slices"
)

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
func (c *Character) ReadAttribute(id AttributeID) interface{} {
	for _, attr := range c.Attributes {
		if attr.ID == id {
			return attr.Value
		}
	}
	return nil // Attribute not found
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
