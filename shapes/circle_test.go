package shapes

import "testing"

func TestCircle_Area(t *testing.T) {
	c := Circle{Radius: 1}
	got := c.Area()
	want := 3.141592653589793
	if got != want {
		t.Errorf("Area() = %v, want %v", got, want)
	}
}

func TestCircle_Circumference(t *testing.T) {
	c := Circle{Radius: 1}
	got := c.Circumference()
	want := 6.283185307179586
	if got != want {
		t.Errorf("Circumference() = %v, want %v", got, want)
	}
}

func TestCircle_Diameter(t *testing.T) {
	c := Circle{Radius: 1}
	got := c.Diameter()
	want := 2.0
	if got != want {
		t.Errorf("Diameter() = %v, want %v", got, want)
	}
}
