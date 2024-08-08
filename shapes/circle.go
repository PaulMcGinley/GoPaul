package shapes

import "math"

type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circumference returns the circumference of the circle
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// Diameter returns the diameter of the circle
func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}
