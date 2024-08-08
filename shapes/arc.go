package shapes

import "math"

type Arc struct {
	Circle Circle
	Angle  float64
}

// SectorArea returns the area of the sector
func (a Arc) SectorArea() float64 {
	return a.Circle.Area() * a.Angle / 360
}

// Chord returns the length of the chord
func (a Arc) Chord() float64 {
	return 2 * a.Circle.Radius * math.Sin(a.Angle*math.Pi/360)
}

// ArcLength returns the length of the arc
func (a Arc) ArcLength() float64 {
	return a.Circle.Radius * a.Angle * math.Pi / 180
}
