package main

type TextStyle struct {
	Font      *Font
	Alignment Alignment
	Borders   []Border
	Boundries *Size
	Padding   Margin
	WordWrap  bool
	Multiline bool
	Overflow  string
}

type Text struct {
	Value string
	Style TextStyle
}

func (t Text) Render(renderer *DocumentRenderer) error {
	return renderer.DrawText(t.Value, &t.Style)
}

func (t Text) GetSize() Size {
	return *t.Style.Boundries
}
