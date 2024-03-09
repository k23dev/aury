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

func (a *Attribute) ToInt() int {
	i, ok := a.Value.(int)

	if ok {
		return i
	}
	return 0
}

func (a *Attribute) ToUint() uint {
	i, ok := a.Value.(uint)

	if ok {
		return i
	}
	return 0
}

func (a *Attribute) ToBool() bool {
	i, ok := a.Value.(bool)

	if ok {
		return i
	}
	return false
}

func (a *Attribute) ToStr() string {
	i, ok := a.Value.(string)

	if ok {
		return i
	}
	return ""
}

func (a *Attribute) ToFloat32() float32 {
	i, ok := a.Value.(float32)

	if ok {
		return i
	}
	return 0.0
}

func (a *Attribute) ToFloat64() float64 {
	i, ok := a.Value.(float64)

	if ok {
		return i
	}
	return 0.0
}
