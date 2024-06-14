package main

type BorderSide int

const (
	Left BorderSide = 1 << iota
	Right
	Top
	Bottom

	AllSides = Left | Right | Top | Bottom
)

func (b BorderSide) IsValid() bool {
	return b&Left != 0 || b&Right != 0 || b&Top != 0 || b&Bottom != 0
}

type Border struct {
	Side    BorderSide
	Options LineOptions
}

func NewBorder() Border {
	return Border{
		Side: AllSides,
	}
}

func NewBorderWithOptions(options LineOptions) Border {
	return Border{
		Side:    AllSides,
		Options: options,
	}
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
