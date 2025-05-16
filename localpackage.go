// The file name is ignored by go
// you can give any file name with the same package name
// the functions would be automatically picked up
// unlike other programming languages the _ is not used for file name separation
// nor the - is used for long names
package main

import (
	"fmt"
	"math"
)

func squared() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
