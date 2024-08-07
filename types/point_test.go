package types

import "testing"

func TestPoint_Abs(t *testing.T) {
	p := Point{X: -1, Y: -1}
	got := p.Abs()
	want := Point{X: 1, Y: 1}
	if got != want {
		t.Errorf("Abs() = %v, want %v", got, want)
	}
}

func TestPoint_Add(t *testing.T) {
	p := Point{X: 1, Y: 1}
	q := Point{X: 2, Y: 2}
	got := p.Add(q)
	want := Point{X: 3, Y: 3}
	if got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}

func TestPoint_Cross(t *testing.T) {
	p := Point{X: 1, Y: 1}
	q := Point{X: 2, Y: 2}
	got := p.Cross(q)
	want := 0
	if got != want {
		t.Errorf("Cross() = %v, want %v", got, want)
	}
}

func TestPoint_Div(t *testing.T) {
	p := Point{X: 4, Y: 4}
	q := Point{X: 2, Y: 2}
	got := p.Div(q)
	want := Point{X: 2, Y: 2}
	if got != want {
		t.Errorf("Div() = %v, want %v", got, want)
	}
}

func TestPoint_Dot(t *testing.T) {
	p := Point{X: 1, Y: 1}
	q := Point{X: 2, Y: 2}
	got := p.Dot(q)
	want := 4
	if got != want {
		t.Errorf("Dot() = %v, want %v", got, want)
	}
}

func TestPoint_Equal(t *testing.T) {
	p := Point{X: 1, Y: 1}
	q := Point{X: 1, Y: 1}
	got := p.Equal(q)
	want := true
	if got != want {
		t.Errorf("Equal() = %v, want %v", got, want)
	}
}

func TestPoint_Mul(t *testing.T) {
	p := Point{X: 2, Y: 2}
	q := Point{X: 2, Y: 2}
	got := p.Mul(q)
	want := Point{X: 4, Y: 4}
	if got != want {
		t.Errorf("Mul() = %v, want %v", got, want)
	}
}

func TestPoint_Sub(t *testing.T) {
	p := Point{X: 2, Y: 2}
	q := Point{X: 1, Y: 1}
	got := p.Sub(q)
	want := Point{X: 1, Y: 1}
	if got != want {
		t.Errorf("Sub() = %v, want %v", got, want)
	}
}

func TestPoint_Neg(t *testing.T) {
	p := Point{X: 1, Y: 1}
	got := p.Neg()
	want := Point{X: -1, Y: -1}
	if got != want {
		t.Errorf("Neg() = %v, want %v", got, want)
	}
}

func TestPoint_String(t *testing.T) {
	p := Point{X: 1, Y: 1}
	got := p.String()
	want := "(1, 1)"
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestPoint_Distance(t *testing.T) {
	p := Point{X: 1, Y: 1}
	q := Point{X: 2, Y: 2}
	got := p.Distance(q)
	want := 1.4142135623730951 // https://www.calculator.net/distance-calculator.html?x11=1&y11=1&x12=2&y12=2&type=1&x=Calculate
	if got != want {
		t.Errorf("Distance() = %v, want %v", got, want)
	}
}

func TestPoint_Split(t *testing.T) {
	p := Point{X: 1, Y: 1}
	gotX, gotY := p.Split()
	wantX, wantY := 1, 1
	if gotX != wantX || gotY != wantY {
		t.Errorf("Split() = (%v, %v), want (%v, %v)", gotX, gotY, wantX, wantY)
	}
}
