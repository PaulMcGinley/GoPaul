package types

import (
	"fmt"
	"math"
)

type Pointf struct {
	X, Y float64
}

// Add returns the sum of two points.
func (p Pointf) Add(q Pointf) Pointf {
	return Pointf{p.X + q.X, p.Y + q.Y}
}

// Sub returns the difference of two points.
func (p Pointf) Sub(q Pointf) Pointf {
	return Pointf{p.X - q.X, p.Y - q.Y}
}

// Mul returns the product of two points.
func (p Pointf) Mul(q Pointf) Pointf {
	return Pointf{p.X * q.X, p.Y * q.Y}
}

// Div returns the division of two points.
func (p Pointf) Div(q Pointf) Pointf {
	return Pointf{p.X / q.X, p.Y / q.Y}
}

// Dot returns the dot product of two points.
func (p Pointf) Dot(q Pointf) float64 {
	return p.X*q.X + p.Y*q.Y
}

// Cross returns the cross product of two points.
func (p Pointf) Cross(q Pointf) float64 {
	return p.X*q.Y - p.Y*q.X
}

// Equal returns true if p is equal to q.
func (p Pointf) Equal(q Pointf) bool {
	return p.X == q.X && p.Y == q.Y
}

// Abs returns the absolute value of a point.
func (p Pointf) Abs() Pointf {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}
	return p
}

// Neg returns the negative of a point.
func (p Pointf) Neg() Pointf {
	return Pointf{-p.X, -p.Y}
}

// String returns a string representation of a point.
func (p Pointf) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

// Distance returns the distance between two points.
func (p Pointf) Distance(q Pointf) float64 {
	dx := q.X - p.X
	dy := q.Y - p.Y
	return math.Sqrt((dx * dx) + (dy * dy))
}

// Split returns the x and y components of a point.
func (p Pointf) Split() (x, y float64) {
	return p.X, p.Y
}
