package objectorientation

import (
	"fmt"
)

// There is no date or datetime data type
type InternalPerson struct {
	firstName string
	lastName  string
	age       int
	gender    string
	email     string
}

func (p *InternalPerson) SetName(firstname string, lastname string) {
	p.firstName = firstname
	p.lastName = lastname
}
func (p *InternalPerson) GetName() string {

	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}
