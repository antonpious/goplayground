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
	students "goplayground/student"
	"goplayground/teachers"
	"math/rand"
)

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

}
