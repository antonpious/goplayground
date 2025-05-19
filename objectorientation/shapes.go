package objectorientation

import (
	"math"
)

// Demonstrate Polymorphism

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// TODO if we convert to int the rounding error needs to be cast
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
