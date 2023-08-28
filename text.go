package main

type Text struct {
	style TextStyle
}

func (t *Text) Render(renderer *DocumentRenderer) error {
	return renderer.Text("Leonardo", &t.style)
}

type TextStyle struct {
	Font      *Font
	Alignment Alignment
	Borders   []Border
	Boundries *Size
	Padding   Margin
	WordWrap  int
}
