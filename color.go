package main

type Color struct {
	R uint8
	G uint8
	B uint8
}

func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) IsEqual(other Color) bool {
	return c.R == other.R && c.G == other.G && c.B == other.B
}
