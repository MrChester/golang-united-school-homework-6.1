package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() (result float64) {
	return 2*r.Height + 2*r.Weight
}

func (r Rectangle) CalcArea() (result float64) {
	return r.Height * r.Weight
}
