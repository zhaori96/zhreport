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

func (s *Size) ToRect() *gopdf.Rect {
	return &gopdf.Rect{
		W: s.Width,
		H: s.Height,
	}
}
