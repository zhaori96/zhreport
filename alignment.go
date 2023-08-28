package main

type Alignment int

const (
	LeftAlignment Alignment = iota + 1<<16
	RightAlignment
	CenterAlignment
	MiddleAlignment
	TopAlignment
	BottomAlignment
)
