package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() float64 {
	var perimeter float64
	perimeter = 2 * (r.Height + r.Weight)
	return perimeter
}

func (r Rectangle) CalcArea() float64 {
	var area float64
	area = r.Height * r.Weight
	return area
}
