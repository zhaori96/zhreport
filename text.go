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
	return renderer.Text(t.Value, &t.Style)
}

type Container struct {
	Size     Size
	Padding  Margin
	Borders  []Border
	Children []Element
}

func (c Container) Render(renderer *DocumentRenderer) error {
	renderer.BoxWithBorders(c.Size, c.Borders...)
	for _, child := range c.Children {
		if err := child.Render(renderer); err != nil {
			return err
		}
	}

	return nil
}
