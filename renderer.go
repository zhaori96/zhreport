package main

import (
	"path"

	"github.com/signintech/gopdf"
)

type rendererState struct {
	Offset      Offset
	StrokeWidth float64
	StrokeColor Color
	Font        Font
	LineStyle   LineStyle
}

type RendererOptions struct {
	PageSize Size
}

type DocumentRenderer struct {
	engine       gopdf.GoPdf
	currentState rendererState
}

func NewDocumentRenderer(options RendererOptions) *DocumentRenderer {
	renderer := &DocumentRenderer{}
	renderer.engine.Start(gopdf.Config{
		Unit:     gopdf.UnitPT,
		PageSize: *options.PageSize.ToRect(),
	})

	renderer.AddFontFamily(standardFontFamily)
	renderer.SetFont(standardFont)

	renderer.currentState.Font = standardFont
	renderer.currentState.Offset = Offset{renderer.GetX(), renderer.GetY()}

	return renderer
}

func (r *DocumentRenderer) StartNewDocument() {
	r.engine.AddPage()
	r.currentState.Offset = r.GetCurrentPosition()
}

func (r *DocumentRenderer) GetCurrentPosition() Offset {
	return Offset{X: r.engine.GetX(), Y: r.engine.GetY()}
}

func (r *DocumentRenderer) GetX() float64 {
	return r.engine.GetX()
}

func (r *DocumentRenderer) GetY() float64 {
	return r.engine.GetY()
}

func (r *DocumentRenderer) AddFontFamily(family FontFamily) error {
	var err error

	if family.HasRegularStyle() {
		filePath := path.Join(family.Path, family.Regular)
		style := gopdf.TtfOption{Style: gopdf.Regular}
		err = r.engine.AddTTFFontWithOption(family.Name, filePath, style)
		if err != nil {
			return err
		}
	}

	if family.HasItalicStyle() {
		filePath := path.Join(family.Path, family.Italic)
		style := gopdf.TtfOption{Style: gopdf.Italic}
		err = r.engine.AddTTFFontWithOption(family.Name, filePath, style)
		if err != nil {
			return err
		}
	}

	if family.HasBoldStyle() {
		filePath := path.Join(family.Path, family.Bold)
		style := gopdf.TtfOption{Style: gopdf.Bold}
		err = r.engine.AddTTFFontWithOption(family.Name, filePath, style)
		if err != nil {
			return err
		}
	}

	if family.HasBoldItalicStyle() {
		filePath := path.Join(family.Path, family.BoldItalic)
		style := gopdf.TtfOption{Style: gopdf.Bold | gopdf.Italic}
		err = r.engine.AddTTFFontWithOption(family.Name, filePath, style)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *DocumentRenderer) AddMultiFontFamilies(families ...FontFamily) error {
	for _, family := range families {
		err := r.AddFontFamily(family)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *DocumentRenderer) SetFont(font Font) error {
	return r.setFont(font, false)
}

func (r *DocumentRenderer) setFont(font Font, keepCurrentState bool) error {
	var family string = r.currentState.Font.Family
	var style int = r.currentState.Font.Style.Combine()
	var size float64 = r.currentState.Font.Size

	if len(font.Family) > 0 {
		family = font.Family
	}

	if font.Size > 0 {
		size = font.Size
	}

	if font.Style != nil {
		style = font.Style.Combine()
	}

	err := r.engine.SetFontWithStyle(family, style, size)

	if err != nil {
		return nil
	}

	if !keepCurrentState {
		r.currentState.Font = font
	}

	return nil
}

func (r *DocumentRenderer) SetFontFamily(family string) error {
	return r.setFontFamily(family, false)
}

func (r *DocumentRenderer) setFontFamily(family string, keepCurrentState bool) error {
	err := r.SetFont(Font{Family: family})
	if err != nil {
		return err
	}

	if !keepCurrentState {
		r.currentState.Font.Family = family
	}

	return nil
}

func (r *DocumentRenderer) SetFontStyle(style FontStyle) error {
	return r.setFontStyle(style, false)
}

func (r *DocumentRenderer) setFontStyle(style FontStyle, keepCurrentState bool) error {
	err := r.engine.SetFontWithStyle(
		r.currentState.Font.Family,
		style.Combine(),
		r.currentState.Font.Size,
	)

	if err != nil {
		return err
	}

	if !keepCurrentState {
		r.currentState.Font.Style = &style
	}

	return nil
}

func (r *DocumentRenderer) SetFontSize(size float64) error {
	return r.setFontSize(size, false)
}

func (r *DocumentRenderer) setFontSize(size float64, keepCurrentState bool) error {
	err := r.engine.SetFontSize(size)
	if err != nil {
		return err
	}

	if !keepCurrentState {
		r.currentState.Font.Size = size
	}

	return nil
}

func (r *DocumentRenderer) SetStrokeWidth(width float64) {
	r.setStrokeWidth(width, false)
}

func (r *DocumentRenderer) setStrokeWidth(width float64, keepCurrentState bool) {
	r.engine.SetLineWidth(width)

	if !keepCurrentState {
		r.currentState.StrokeWidth = width
	}
}

func (r *DocumentRenderer) SetStrokeColor(color Color) {
	r.setStrokeColor(color, false)

}

func (r *DocumentRenderer) setStrokeColor(color Color, keepCurrentState bool) {
	r.engine.SetStrokeColor(color.R, color.G, color.B)

	if !keepCurrentState {
		r.currentState.StrokeColor = color
	}
}

func (r *DocumentRenderer) SetLineStyle(style LineStyle) {
	r.setLineStyle(style, false)
}

func (r *DocumentRenderer) setLineStyle(style LineStyle, keepCurrentState bool) {
	r.engine.SetLineType(string(style))

	if !keepCurrentState {
		r.currentState.LineStyle = style
	}
}

func (r *DocumentRenderer) Text(text string, style *TextStyle) error {
	if style == nil {
		return r.engine.Cell(nil, text)
	}

	if style.Font != nil {
		//TODO: Check if r.SetFont really need to receive a pointer of Font.
		r.setFont(*style.Font, true)
		defer r.SetFont(r.currentState.Font)
	}

	for _, border := range style.Borders {
		if border.Side == 0 {
			continue
		}

		if border.Side&Left != 0 {
			r.VerticalLine(style.Boundries.Height, &border.Options)
		}

		if border.Side&Right != 0 {
			offset := NewOffset(style.Boundries.Width, r.GetY())
			r.VerticalLineWithOffset(style.Boundries.Height, offset, &border.Options)
		}

		if border.Side&Top != 0 {
			r.HorizontalLine(style.Boundries.Width, &border.Options)
		}

		if border.Side&Bottom != 0 {
			offset := NewOffset(r.GetX(), style.Boundries.Height)
			r.HorizontalLineWithOffset(style.Boundries.Width, offset, &border.Options)
		}

	}

	//TODO: Implement the Padding of style.

	r.engine.CellWithOption(
		&gopdf.Rect{
			W: style.Boundries.Width,
			H: style.Boundries.Height,
		},
		text,
		gopdf.CellOption{
			Align: int(style.Alignment),
		},
	)

	return nil
}

func (r *DocumentRenderer) Line(size Size, offset Offset, options *LineOptions) {
	if options != nil {
		if options.StrokeWidth != r.currentState.StrokeWidth {
			r.setStrokeWidth(options.StrokeWidth, true)
			defer r.SetStrokeWidth(r.currentState.StrokeWidth)
		}

		if !r.currentState.StrokeColor.IsEqual(*options.Color) {
			r.setStrokeColor(*options.Color, true)
			defer r.SetStrokeColor(r.currentState.StrokeColor)
		}
	}

	if size.Width == 0 {
		size.Width = r.GetX()
	}

	if size.Height == 0 {
		size.Height = r.GetY()
	}

	r.setLineStyle(options.Style, true)
	defer r.SetLineStyle(r.currentState.LineStyle)

	r.engine.Line(offset.X, offset.Y, size.Width, size.Height)
}

func (r *DocumentRenderer) HorizontalLine(width float64, options *LineOptions) {
	r.Line(NewSize(width, 0), r.GetCurrentPosition(), options)
}

func (r *DocumentRenderer) HorizontalLineWithOffset(width float64, offset Offset, options *LineOptions) {
	r.Line(NewSize(width, offset.Y), offset, options)
}

func (r *DocumentRenderer) VerticalLine(height float64, options *LineOptions) {
	r.Line(NewSize(0, height), r.GetCurrentPosition(), options)
}

func (r *DocumentRenderer) VerticalLineWithOffset(height float64, offset Offset, options *LineOptions) {
	r.Line(NewSize(offset.X, height), offset, options)
}
