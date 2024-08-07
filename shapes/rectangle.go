package shapes

import (
	"fmt"
	"github.com/PaulMcGinley/GoPaul/types"
	"math"
)

type Rectangle struct {
	Width, Height float64
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Diagonal returns the diagonal of a rectangle.
func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

// String returns a string representation of a rectangle.
func (r Rectangle) String() string {
	return fmt.Sprintf("(%f, %f)", r.Width, r.Height)
}

// Less returns true if r area is less than s area.
func (r Rectangle) Less(s Rectangle) bool {
	return r.Area() < s.Area()
}

// Equal returns true if r area is equal to s area.
func (r Rectangle) Equal(s Rectangle) bool {
	return r.Area() == s.Area()
}

// Greater returns true if r area is greater than s area.
func (r Rectangle) Greater(s Rectangle) bool {
	return r.Area() > s.Area()
}

// LessEqual returns true if r area is less than or equal to s area.
func (r Rectangle) LessEqual(s Rectangle) bool {
	return r.Less(s) || r.Equal(s)
}

// GreaterEqual returns true if r area is greater than or equal to s area.
func (r Rectangle) GreaterEqual(s Rectangle) bool {
	return r.Greater(s) || r.Equal(s)
}

// Min returns the smaller of two rectangles.
func (r Rectangle) Min(s Rectangle) Rectangle {
	if r.Less(s) {
		return r
	}
	return s
}

// Max returns the larger of two rectangles.
func (r Rectangle) Max(s Rectangle) Rectangle {
	if r.Less(s) {
		return s
	}
	return r
}

// Contains returns true if a point is inside a rectangle.
func (r Rectangle) Contains(p types.Point) bool {
	return p.X >= 0 && p.X <= int(r.Width) && p.Y >= 0 && p.Y <= int(r.Height)
}

// Containsf returns true if a pointf is inside a rectangle.
func (r Rectangle) Containsf(p types.Pointf) bool {
	return p.X >= 0 && p.X <= r.Width && p.Y >= 0 && p.Y <= r.Height
}
