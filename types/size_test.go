package types

import "testing"

func TestSize_Abs(t *testing.T) {
	s := Size{Width: -1, Height: -1}
	got := s.Abs()
	want := Size{Width: 1, Height: 1}
	if got != want {
		t.Errorf("Abs() = %v, want %v", got, want)
	}
}

func TestSize_Add(t *testing.T) {
	s := Size{Width: 1, Height: 1}
	q := Size{Width: 2, Height: 2}
	got := s.Add(q)
	want := Size{Width: 3, Height: 3}
	if got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}

func TestSize_Area(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	got := s.Area()
	want := 6
	if got != want {
		t.Errorf("Area() = %v, want %v", got, want)
	}
}

func TestSize_Perimeter(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	got := s.Perimeter()
	want := 10
	if got != want {
		t.Errorf("Perimeter() = %v, want %v", got, want)
	}
}

func TestSize_Equal(t *testing.T) {
	s := Size{Width: 1, Height: 1}
	q := Size{Width: 1, Height: 1}
	got := s.Equal(q)
	want := true
	if got != want {
		t.Errorf("Equal() = %v, want %v", got, want)
	}
}

func TestSize_Less(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	q := Size{Width: 3, Height: 2}
	got := s.Less(q)
	want := false
	if got != want {
		t.Errorf("Less() = %v, want %v", got, want)
	}
}

func TestSize_LessEqual(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	q := Size{Width: 3, Height: 2}
	got := s.LessEqual(q)
	want := true
	if got != want {
		t.Errorf("LessEqual() = %v, want %v", got, want)
	}
}

func TestSize_Greater(t *testing.T) {
	s := Size{Width: 3, Height: 2}
	q := Size{Width: 2, Height: 3}
	got := s.Greater(q)
	want := false
	if got != want {
		t.Errorf("Greater() = %v, want %v", got, want)
	}
}

func TestSize_GreaterEqual(t *testing.T) {
	s := Size{Width: 3, Height: 2}
	q := Size{Width: 2, Height: 3}
	got := s.GreaterEqual(q)
	want := true
	if got != want {
		t.Errorf("GreaterEqual() = %v, want %v", got, want)
	}
}

func TestSize_Mul(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	i := 2
	got := s.Mul(i)
	want := 12
	if got != want {
		t.Errorf("Mul() = %v, want %v", got, want)
	}
}

func TestSize_Mulf(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	f := 2.0
	got := s.Mulf(f)
	want := 12.0
	if got != want {
		t.Errorf("Mulf() = %v, want %v", got, want)
	}
}

func TestSize_Sub(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	q := Size{Width: 1, Height: 2}
	got := s.Sub(q)
	want := Size{Width: 1, Height: 1}
	if got != want {
		t.Errorf("Sub() = %v, want %v", got, want)
	}
}

func TestSize_Div(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	i := 2
	got := s.Div(i)
	want := 3
	if got != want {
		t.Errorf("Div() = %v, want %v", got, want)
	}
}

func TestSize_Divf(t *testing.T) {
	s := Size{Width: 2, Height: 3}
	f := 2.0
	got := s.Divf(f)
	want := 3.0
	if got != want {
		t.Errorf("Divf() = %v, want %v", got, want)
	}
}
