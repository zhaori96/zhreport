package main

type LineStyle string

const (
	LineStyleDotted LineStyle = "dotted"
	LineStyleDashed LineStyle = "dashed"
	LineStyleSolid  LineStyle = "solid"
)

type LineOptions struct {
	StrokeWidth float64
	Style       LineStyle
	Color       *Color
}
