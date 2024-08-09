package types

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

// Add returns the sum of two points.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub returns the difference of two points.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Mul returns the product of two points.
func (p Point) Mul(q Point) Point {
	return Point{p.X * q.X, p.Y * q.Y}
}

// Div returns the division of two points.
func (p Point) Div(q Point) Point {
	return Point{p.X / q.X, p.Y / q.Y}
}

// Dot returns the dot product of two points.
func (p Point) Dot(q Point) int {
	return p.X*q.X + p.Y*q.Y
}

// Cross returns the cross product of two points.
func (p Point) Cross(q Point) int {
	return p.X*q.Y - p.Y*q.X
}

// Equal returns true if p is equal to q.
func (p Point) Equal(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

// Abs returns the absolute value of a point.
func (p Point) Abs() Point {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}
	return p
}

// Neg returns the negative of a point.
func (p Point) Neg() Point {
	return Point{-p.X, -p.Y}
}

// String returns the string representation of a point.
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Parse returns a point from a string.
func (p Point) Parse(s string) (Point, error) {
	var x, y int
	_, err := fmt.Sscanf(s, "(%d, %d)", &x, &y)
	if err != nil {
		return Point{}, err
	}
	return Point{x, y}, nil
}

func (p Point) TryParse(s string) (Point, bool) {
	var x, y int
	format := "(%d,%d)"
	// Todo: remove spaces
	_, err := fmt.Sscanf(s, format, &x, &y)
	if err != nil {
		return Point{}, false
	}
	return Point{x, y}, true
}

// ManhattanDistance returns the Manhattan distance between two points.
func (p Point) ManhattanDistance(q Point) int {
	return int(math.Abs(float64(p.X-q.X)) + math.Abs(float64(p.Y-q.Y)))
}

// EuclideanDistance returns the Euclidean distance between two points.
func (p Point) EuclideanDistance(q Point) float64 {
	return math.Sqrt(float64((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)))
}

// Distance returns the distance between two points.
func (p Point) Distance(q Point) float64 {
	dx := q.X - p.X
	dy := q.Y - p.Y
	return math.Sqrt(float64((dx * dx) + (dy * dy)))
}

// Split returns the x and y coordinates of a point as two integers.
func (p Point) Split() (x, y int) {
	return p.X, p.Y
}
