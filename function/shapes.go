package function

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
