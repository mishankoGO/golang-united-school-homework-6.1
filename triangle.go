package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	var perimeter float64
	perimeter = 3 * t.Side
	return perimeter
}

func (t Triangle) CalcArea() float64 {
	var area float64
	area = math.Sqrt(3) / 4 * t.Side * t.Side
	return area
}
