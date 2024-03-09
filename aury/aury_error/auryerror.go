package auryError

import (
	"fmt"
)

type AuryError struct {
	Name    string
	Message string
	Code    uint
}

func New(name, msg string, code uint) error {
	aerr := AuryError{
		Name:    name,
		Message: msg,
		Code:    code,
	}
	return &aerr
}

func (aerr *AuryError) Error() string {
	return fmt.Sprintf("name %s: (CODE: %d) err %s", aerr.Name, aerr.Code, aerr.Message)
}
