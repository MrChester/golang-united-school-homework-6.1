package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) (err error) {

	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return fmt.Errorf("added more than max")
	}

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}

	extractedShape := b.shapes[i]
	tempSlice := append(b.shapes[:i], b.shapes[i+1:]...)
	b.shapes = tempSlice

	return extractedShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}
	
	tempShape := b.shapes[i]
	b.shapes[i] = shape

	return tempShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (output float64) {
	for _, shape := range b.shapes {
		output += shape.CalcPerimeter()
	}
	
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (output float64) {
	for _, shape := range b.shapes {
		output += shape.CalcArea()
	}

	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var tempSlice []Shape
	var circles int = 0
	
	for _, shape := range b.shapes {
		if _, ok := shape.(*Circle); !ok {
			tempSlice = append(tempSlice, shape)
		} else {
			circles =+  1
		}
		
	}

	if circles == 0 {
		return fmt.Errorf("no circles")
	}

	b.shapes = tempSlice

	return nil
}
