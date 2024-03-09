package characters

// Attribute represents an attribute with a name, type, and value.
type Attribute struct {
	ID    AttributeID // ID of the attribute
	Type  string      // Type of the attribute
	Value interface{} // Value of the attribute (generic type)
}

type AttributeID string

// AddAttribute adds a new attribute to the provided slice.
func New(id AttributeID, typ string, value interface{}) *Attribute {
	return &Attribute{
		ID:    id,
		Type:  typ,
		Value: value,
	}
}
