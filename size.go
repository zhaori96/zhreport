package main

import (
	"fmt"

	"github.com/signintech/gopdf"
)

type Size struct {
	Width  float64
	Height float64
}

func NewSize(width, height float64) Size {
	return Size{
		Width:  width,
		Height: height,
	}
}

func NewSizeWithPadding(width, height float64, padding Margin) Size {
	size := Size{
		Width:  width,
		Height: height,
	}

	return size.WithPadding(padding)
}

func (s Size) IsZero() bool {
	return s.Width == 0 && s.Height == 0
}

func (s *Size) ToRect() *gopdf.Rect {
	return &gopdf.Rect{
		W: s.Width,
		H: s.Height,
	}
}

func (s Size) ToOffset() Offset {
	return Offset{X: s.Width, Y: s.Height}
}

func (s Size) WithPadding(padding Margin) Size {
	return Size{
		Width:  s.Width - padding.Left - padding.Right,
		Height: s.Height - padding.Top - padding.Bottom,
	}
}

func (s Size) WithHorizontalPadding(left float64, right float64) Size {
	return Size{
		Width:  s.Width - left - right,
		Height: s.Height,
	}
}

func (s Size) WithVerticalPadding(top float64, bottom float64) Size {
	return Size{
		Width:  s.Width,
		Height: s.Height - top - bottom,
	}
}

func (s Size) FromAxis(axis Axis) float64 {
	switch axis {
	case HorizontalAxis:
		return s.Width
	case VerticalAxis:
		return s.Height
	default:
		panic(fmt.Errorf("Size.FromAxis: %w", ErrInvalidAxis))
	}
}
