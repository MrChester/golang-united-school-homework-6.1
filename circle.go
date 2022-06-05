package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() (result float64) {
	return 2 * math.Pi * c.Radius
}


func (c Circle) CalcArea() (result float64) {
	return math.Pi * c.Radius * c.Radius
}