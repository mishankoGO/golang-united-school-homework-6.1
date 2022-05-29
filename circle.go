package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	var perimeter float64
	perimeter = 2 * c.Radius * math.Pi
	return perimeter
}

func (c Circle) CalcArea() float64 {
	var area float64
	area = c.Radius * c.Radius * math.Pi
	return area
}
