package main

type BorderSide int

const (
	Left     BorderSide = 1
	Right    BorderSide = 2
	Top      BorderSide = 4
	Bottom   BorderSide = 8
	AllSides BorderSide = Left | Right | Top | Bottom
)

type Border struct {
	Side    BorderSide
	Options LineOptions
}

func NewBorderLeft() Border {
	return Border{
		Side:    Left,
		Options: LineOptions{},
	}
}

func NewBorderLeftWithOptions(options LineOptions) Border {
	return Border{
		Side:    Left,
		Options: options,
	}
}

func NewBorderRight() Border {
	return Border{
		Side:    Right,
		Options: LineOptions{},
	}
}

func NewBorderRightWithOptions(options LineOptions) Border {
	return Border{
		Side:    Right,
		Options: options,
	}
}

func NewBorderTop() Border {
	return Border{
		Side:    Top,
		Options: LineOptions{},
	}
}

func NewBorderTopWithOptions(options LineOptions) Border {
	return Border{
		Side:    Top,
		Options: options,
	}
}

func NewBorderBottom() Border {
	return Border{
		Side:    Bottom,
		Options: LineOptions{},
	}
}

func NewBorderBottomWithOptions(options LineOptions) Border {
	return Border{
		Side:    Bottom,
		Options: options,
	}
}
