package main

type Margin struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func NewMargin(left, right, top, bottom float64) Margin {
	return Margin{
		Left:   left,
		Right:  right,
		Top:    top,
		Bottom: bottom,
	}
}

func NewHorizontalMargin(left, right float64) Margin {
	return Margin{
		Left:  left,
		Right: right,
	}
}

func NewVerticalMargin(top, bottom float64) Margin {
	return Margin{
		Top:   top,
		Bottom: bottom,
	}
}
