package types

type Size struct {
	Width, Height int
}

func (s Size) Area() int {
	return s.Width * s.Height
}

func (s Size) Perimeter() int {
	return 2 * (s.Width + s.Height)
}

func (s Size) Equal(q Size) bool {
	return s.Area() == q.Area()
}

func (s Size) Less(q Size) bool {
	return s.Area() < q.Area()
}

func (s Size) LessEqual(q Size) bool {
	return s.Less(q) || s.Equal(q)
}

func (s Size) Greater(q Size) bool {
	return s.Area() > q.Area()
}

func (s Size) GreaterEqual(q Size) bool {
	return s.Greater(q) || s.Equal(q)
}

func (s Size) Abs() Size {
	if s.Width < 0 {
		s.Width = -s.Width
	}
	if s.Height < 0 {
		s.Height = -s.Height
	}
	return s
}

func (s Size) Add(q Size) Size {
	return Size{s.Width + q.Width, s.Height + q.Height}
}

func (s Size) Sub(q Size) Size {
	return Size{s.Width - q.Width, s.Height - q.Height}
}

func (s Size) Mul(i int) int {
	return s.Width * i * s.Height
}

func (s Size) Mulf(f float64) float64 {
	return float64(s.Width) * f * float64(s.Height)
}

func (s Size) Div(i int) int {
	return s.Width / i * s.Height
}

func (s Size) Divf(f float64) float64 {
	return float64(s.Width) / f * float64(s.Height)
}
