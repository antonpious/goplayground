package main

/*
Imports can be together or single

When importing other packages
The folder name is the import
The file name can be inside that import
e.g student is the folder name but students in the package
if this is the case an alias is required
The file and folder name is the same
then the import can be just the folder
e.g teachers is the folder and file name
*/
import (
	"fmt"
	"goplayground/objectorientation"
	students "goplayground/student"
	"goplayground/teachers"
	"math/rand"
)

// Have to find a way to import specific types
// Like the person type persent here
// "goplayground/objectorientation"
// Note: The VS code does an automatic linting
// so you might see your code getting lost when saving
// while learning you can disable it

// Only one main function is allowed across packages
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// This would be called as its the same package
	// no import required
	squared()
	// This would call as separate package
	// The folder name is different from package name
	students.GetStudents()
	// The folder name is same as the package name
	teachers.GetTeachers()

	// object orientation
	// note the special syntax. there is a : and = to assign this variable
	person1 := objectorientation.Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Gender:    "Male",
		Email:     "johndoe@gmailtest.com"}

	fmt.Println("Full Name ", person1.GetName())

	person1.SetName("Smith", "Brown")

	fmt.Println("Full Name ", person1.GetName())

	// There is no real encapsulation
	// one can directly update the struct if the struct is exposed as
	// public variables
	person1.FirstName = "Peter"
	fmt.Println("Full Name ", person1.GetName())

	// There is encapsulation
	// if the fields are not exposed and only available through functions
	person2 := objectorientation.InternalPerson{}

	// This should print an empty name
	fmt.Println("Full Name ", person2.GetName())

	person2.SetName("Jane", "Smith")

	fmt.Println("Full Name ", person2.GetName())

}
