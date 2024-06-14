package main

type Alignment int

const (
	LeftAlignment   Alignment = 8
	RightAlignment  Alignment = 2
	CenterAlignment Alignment = 16
	MiddleAlignment Alignment = 32
	TopAlignment    Alignment = 4
	BottomAlignment Alignment = 1
)

func (a Alignment) IsValid() bool {
	switch a {
	case LeftAlignment, RightAlignment,
		CenterAlignment, MiddleAlignment,
		TopAlignment, BottomAlignment:
		return true
	default:
		return false
	}
}
