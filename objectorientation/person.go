package objectorientation

import (
	"fmt"
)

// There is no date or datetime data type
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Email     string
}

func (p *Person) SetName(firstname string, lastname string) {
	p.FirstName = firstname
	p.LastName = lastname
}
func (p *Person) GetName() string {

	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
