package characters

type Attribute struct {
	Name  string
	Value interface{}
}

func NewAttribute(name string, value interface{}) *Attribute {
	return &Attribute{}
}
