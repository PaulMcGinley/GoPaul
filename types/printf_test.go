package types

import "testing"

func TestPointf_Abs(t *testing.T) {
	p := Pointf{X: -1, Y: -1}
	got := p.Abs()
	want := Pointf{X: 1, Y: 1}
	if got != want {
		t.Errorf("Abs() = %v, want %v", got, want)
	}
}

func TestPointf_Add(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	q := Pointf{X: 2, Y: 2}
	got := p.Add(q)
	want := Pointf{X: 3, Y: 3}
	if got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}

func TestPointf_Cross(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	q := Pointf{X: 2, Y: 2}
	got := p.Cross(q)
	var want = float64(0)
	if got != want {
		t.Errorf("Cross() = %v, want %v", got, want)
	}
}

func TestPointf_Div(t *testing.T) {
	p := Pointf{X: 4, Y: 4}
	q := Pointf{X: 2, Y: 2}
	got := p.Div(q)
	want := Pointf{X: 2, Y: 2}
	if got != want {
		t.Errorf("Div() = %v, want %v", got, want)
	}
}

func TestPointf_Dot(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	q := Pointf{X: 2, Y: 2}
	got := p.Dot(q)
	var want = float64(4)
	if got != want {
		t.Errorf("Dot() = %v, want %v", got, want)
	}
}

func TestPointf_Equal(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	q := Pointf{X: 1, Y: 1}
	got := p.Equal(q)
	want := true
	if got != want {
		t.Errorf("Equal() = %v, want %v", got, want)
	}
}

func TestPointf_Mul(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	q := Pointf{X: 2, Y: 2}
	got := p.Mul(q)
	want := Pointf{X: 2, Y: 2}
	if got != want {
		t.Errorf("Mul() = %v, want %v", got, want)
	}
}

func TestPointf_Neg(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	got := p.Neg()
	want := Pointf{X: -1, Y: -1}
	if got != want {
		t.Errorf("Neg() = %v, want %v", got, want)
	}
}

func TestPointf_Sub(t *testing.T) {
	p := Pointf{X: 2, Y: 2}
	q := Pointf{X: 1, Y: 1}
	got := p.Sub(q)
	want := Pointf{X: 1, Y: 1}
	if got != want {
		t.Errorf("Sub() = %v, want %v", got, want)
	}
}

func TestPointf_String(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	got := p.String()
	want := "(1.000000, 1.000000)"
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestPointf_Distance(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	q := Pointf{X: 2, Y: 2}
	got := p.Distance(q)
	want := 1.4142135623730951
	if got != want {
		t.Errorf("Distance() = %v, want %v", got, want)
	}
}

func TestPointf_Split(t *testing.T) {
	p := Pointf{X: 1, Y: 1}
	got1, got2 := p.Split()
	want1, want2 := 1.0, 1.0
	if got1 != want1 || got2 != want2 {
		t.Errorf("Split() = (%v, %v), want (%v, %v)", got1, got2, want1, want2)
	}
}
