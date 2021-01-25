package interfaces

import "math"

// Shape ...
type Shape interface {
	Area() float64
}

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle ...
type Circle struct {
	Radius float64
}

// Triangle ...
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter takes width and height than returns perimeter
func Perimeter(rectangle Rectangle) float64 { //function
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area ...
func (r Rectangle) Area() float64 { //method
	return r.Width * r.Height
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area ...
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
