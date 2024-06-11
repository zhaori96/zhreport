package main

import "github.com/signintech/gopdf"

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

func (s *Size) ToRect() *gopdf.Rect {
	return &gopdf.Rect{
		W: s.Width,
		H: s.Height,
	}
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
