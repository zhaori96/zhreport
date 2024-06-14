package main

import "slices"

type Element interface {
	GetSize() Size
	Render(renderer *DocumentRenderer) error
}

type Elements []Element

func (e Elements) TotalWidth() float64 {
	total := 0.0
	for _, element := range e {
		total += element.GetSize().Width
	}
	return total
}

func (e Elements) MaxWitdh() float64 {
	return slices.MaxFunc(e, func(a, b Element) int {
		return int(a.GetSize().Width) - int(b.GetSize().Width)
	}).GetSize().Width
}

func (e Elements) MaxHeight() float64 {
	return slices.MaxFunc(e, func(a, b Element) int {
		return int(a.GetSize().Height) - int(b.GetSize().Height)
	}).GetSize().Height
}

func (e Elements) TotalHeight() float64 {
	total := 0.0
	for _, element := range e {
		total += element.GetSize().Height
	}
	return total
}

func (e Elements) TotalBoundries() Size {
	total := Size{}
	for _, element := range e {
		boundries := element.GetSize()
		total.Width += boundries.Width
		total.Height += boundries.Height
	}
	return total
}

func (e Elements) TotalFromAxis(axis Axis) float64 {
	switch axis {
	case HorizontalAxis:
		return e.TotalWidth()
	case VerticalAxis:
		return e.TotalHeight()
	default:
		panic("Elements.TotalFromAxis: invalid axis")
	}
}

func (e Elements) MaxFromAxis(axis Axis) float64 {
	switch axis {
	case HorizontalAxis:
		return e.MaxWitdh()
	case VerticalAxis:
		return e.MaxHeight()
	default:
		panic("Elements.MaxFromAxis: invalid axis")
	}
}

func (e Elements) FitsParent(parent Size, axis Axis) bool {
	switch axis {
	case HorizontalAxis:
		return parent.Width >= e.TotalWidth() && parent.Height >= e.MaxHeight()
	case VerticalAxis:
		return parent.Height >= e.TotalHeight() && parent.Width >= e.MaxWitdh()
	default:
		panic("List.Render: invalid Axis")
	}
}

func (e Elements) FitsParentWithPadding(
	parent Size,
	padding Margin,
	axis Axis,
) bool {
	return e.FitsParent(parent.WithPadding(padding), axis)
}

func (e Elements) SpaceBetween(parent Size, axis Axis) Offset {
	if parent.IsZero() {
		return NewZeroOffset()
	}

	elementCount := len(e)
	if elementCount == 0 {
		return NewZeroOffset()
	}

	switch axis {
	case HorizontalAxis:
		totalSize := e.TotalWidth()
		return NewOffsetX((parent.Width - totalSize) / float64(elementCount-1))
	case VerticalAxis:
		totalSize := e.TotalHeight()
		return NewOffsetY((parent.Height - totalSize) / float64(elementCount-1))
	default:
		panic("Elements.Between: invalid axis")
	}
}

func (e Elements) SpaceEvenly(parent Size, axis Axis) Offset {
	if parent.IsZero() {
		return NewZeroOffset()
	}

	elementCount := len(e)
	if elementCount == 0 {
		return NewZeroOffset()
	}

	switch axis {
	case HorizontalAxis:
		totalSize := e.TotalWidth()
		return NewOffsetX((parent.Width - totalSize) / float64(elementCount+1))
	case VerticalAxis:
		totalSize := e.TotalHeight()
		return NewOffsetY((parent.Height - totalSize) / float64(elementCount+1))
	default:
		panic("Elements.SpaceEvenly: invalid axis")
	}
}

func (e Elements) SpaceAround(parent Size, axis Axis) (Offset, Offset) {
	if parent.IsZero() {
		return NewZeroOffset(), NewZeroOffset()
	}

	elementCount := len(e)
	if elementCount == 0 {
		return NewZeroOffset(), NewZeroOffset()
	}

	switch axis {
	case HorizontalAxis:
		totalSize := e.TotalWidth()
		extraSpace := parent.Width - totalSize
		spacing := extraSpace / float64(elementCount*2)
		return NewOffsetX(spacing), NewOffsetX(spacing * 2)
	case VerticalAxis:
		totalSize := e.TotalHeight()
		extraSpace := parent.Height - totalSize
		spacing := extraSpace / float64(elementCount*2)
		return NewOffsetY(spacing), NewOffsetY(spacing * 2)

	default:
		panic("Elements.SpaceAround: invalid axis")
	}
}

func (e Elements) CalculateSpacing(
	parent Size,
	justify JustifyContent,
	axis Axis,
) (Offset, Offset) {
	var edgeGap, betweenGap Offset
	switch justify {
	case JustifyContentSpaceAround:
		edgeGap, betweenGap = e.SpaceAround(parent, axis)

	case JustifyContentSpaceBetween:
		betweenGap = e.SpaceBetween(parent, axis)

	case JustifyContentSpaceEvenly:
		betweenGap = e.SpaceEvenly(parent, axis)
		edgeGap = betweenGap
	}
	return edgeGap, betweenGap
}

func (e Elements) CalculateSpacingWithPadding(
	parent Size,
	padding Margin,
	justify JustifyContent,
	axis Axis,
) (Offset, Offset) {
	return e.CalculateSpacing(parent.WithPadding(padding), justify, axis)
}
