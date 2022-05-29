package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorShapesQuantity = errors.New("shapes quantity goes out of shapesCapacity")
	errorInvalidIndex   = errors.New("invalid index")
	errorNoCircles      = errors.New("no circles")
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
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity > len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return fmt.Errorf("%w", errorShapesQuantity)
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, fmt.Errorf("%w", errorInvalidIndex)
	} else {
		return b.shapes[i], nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, fmt.Errorf("%w", errorInvalidIndex)
	} else {
		var res Shape
		res = b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return res, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, fmt.Errorf("%w", errorInvalidIndex)
	} else {
		var res Shape
		res = b.shapes[i]
		b.shapes[i] = shape
		return res, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64
	for _, elem := range b.shapes {
		sumPerimeter += elem.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for _, elem := range b.shapes {
		sumArea += elem.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var indexes []int
	for i, elem := range b.shapes {
		if _, ok := (elem).(*Circle); ok {
			indexes = append(indexes, i)
		}
	}

	if len(indexes) == 0 {
		return fmt.Errorf("%w", errorNoCircles)
	} else {
		for i := len(indexes) - 1; i >= 0; i-- {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		}
		return nil
	}
}
