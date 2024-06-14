package main

type Axis int

const (
	HorizontalAxis = iota
	VerticalAxis
)

func (a Axis) IsValid() bool {
	return a >= 0 && a <= 1
}
