package golang_united_school_homework

import "fmt"

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
		return fmt.Errorf("%s", "shapes quantity goes out of shapesCapacity")
	}
	//panic("implement me")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, fmt.Errorf("%s", "invalid index")
	} else {
		return b.shapes[i], nil
	}
	//panic("implement me")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, fmt.Errorf("%s", "invalid index")
	} else {
		var res Shape
		res = b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return res, nil
	}
	//panic("implement me")

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i > len(b.shapes)-1 {
		return nil, fmt.Errorf("%s", "invalid index")
	} else {
		var res Shape
		res = b.shapes[i]
		b.shapes[i] = shape
		return res, nil
	}
	//panic("implement me")

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64
	for _, elem := range b.shapes {
		sumPerimeter += elem.CalcPerimeter()
	}
	return sumPerimeter
	//panic("implement me")

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for _, elem := range b.shapes {
		sumArea += elem.CalcArea()
	}
	return sumArea
	//panic("implement me")

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var NoCircles []Shape
	for _, elem := range b.shapes {
		if shape, ok := elem.(Circle); !ok {
			fmt.Println(shape, &shape)
			NoCircles = append(NoCircles, &shape)
		}
	}
	fmt.Println(len(NoCircles), len(b.shapes))
	fmt.Println(NoCircles, b.shapes)
	if len(NoCircles) == len(b.shapes) {
		return fmt.Errorf("%s", "No circles in the shapes")
	} else {
		b.shapes = NoCircles
		fmt.Println(b.shapes)

		return nil
	}
	//panic("implement me")

}
