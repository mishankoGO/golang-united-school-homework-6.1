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
	semiPerimeter := t.Side * 3 / 2
	area = math.Sqrt(semiPerimeter * (semiPerimeter - t.Side) * (semiPerimeter - t.Side) * (semiPerimeter - t.Side))
	return area
}
