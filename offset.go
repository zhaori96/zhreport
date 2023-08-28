package main

type Offset struct {
	X float64
	Y float64
}

func NewOffset(x float64, y float64) Offset {
	return Offset{
		X: x,
		Y: y,
	}
}
