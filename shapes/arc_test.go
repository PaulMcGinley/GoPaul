package shapes

import "testing"

func TestArc_ArcLength(t *testing.T) {
	c := Circle{Radius: 10}
	a := Arc{Circle: c, Angle: 90}
	want := 15.707963267948966
	got := a.ArcLength()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestArc_Chord(t *testing.T) {
	c := Circle{Radius: 10}
	a := Arc{Circle: c, Angle: 90}
	want := 14.14213562373095
	got := a.Chord()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestArc_SectorArea(t *testing.T) {
	c := Circle{Radius: 10}
	a := Arc{Circle: c, Angle: 90}
	want := 78.53981633974483
	got := a.SectorArea()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
