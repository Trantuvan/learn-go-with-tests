package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

type Circle struct {
	Radius float64
}

func (cir Circle) Perimeter() float64 {
	return cir.Radius * 2 * math.Pi
}

func (cir Circle) Area() float64 {
	return math.Pi * math.Pow(cir.Radius, 2)
}
