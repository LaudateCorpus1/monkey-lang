package object

import (
	"fmt"
)

// String is the string type used to represent string literals and holds
// an internal string value
type String struct {
	Value string
}

func (s *String) Equal(other Object) bool {
	if obj, ok := other.(*String); ok {
		return s.Value == obj.Value
	}
	return false
}

func (s *String) String() string {
	return s.Value
}

// Clone creates a new copy
func (s *String) Clone() Object {
	return &String{Value: s.Value}
}

// Type returns the type of the object
func (s *String) Type() Type { return STRING }

// Inspect returns a stringified version of the object for debugging
func (s *String) Inspect() string { return fmt.Sprintf("%#v", s.Value) }
