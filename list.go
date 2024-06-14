package main

type List struct {
	Axis          Axis
	MainAxisSize  float64
	CrossAxisSize float64
	Padding       Margin
	Borders       []Border
	Justify       JustifyContent
	Children      Elements

	size Size
}

func (l List) GetSize() Size {
	return l.size
}

func (l List) Render(renderer *DocumentRenderer) error {
	if !l.Axis.IsValid() {
		return ErrInvalidAxis
	}

	l.size = l.computeSize()
	if l.size.IsZero() {
		return ErrInvalidSize
	}

	paddedSize := l.size.WithPadding(l.Padding)
	if !l.Children.FitsParent(paddedSize, l.Axis) {
		return ErrElementOverflow
	}

	defer renderer.SetOffset(renderer.GetCurrentOffset())

	err := renderer.DrawBoxWithBorders(l.size, l.Borders...)
	if err != nil {
		return err
	}
	renderer.AddXY(l.Padding.Left, l.Padding.Top)

	edgeGap, gap := l.Children.CalculateSpacing(paddedSize, l.Justify, l.Axis)
	renderer.AddOffset(edgeGap)

	for index, child := range l.Children {
		if err := child.Render(renderer); err != nil {
			return err
		}

		renderer.AddAxisOffset(child.GetSize().ToOffset(), l.Axis)
		if index < len(l.Children)-1 {
			renderer.AddOffset(gap)
		}
	}

	return nil
}

func (l *List) computeSize() Size {
	switch l.Axis {
	case HorizontalAxis:
		return Size{Width: l.MainAxisSize, Height: l.CrossAxisSize}
	case VerticalAxis:
		return Size{Width: l.CrossAxisSize, Height: l.MainAxisSize}
	default:
		return Size{}
	}
}
