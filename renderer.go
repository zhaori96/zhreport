package main

import (
	"math"
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
	PageSize             Size
	Padding              Margin
	DefaultSeparatorSize float64
}

type Element interface {
	Render(renderer *DocumentRenderer) error
}

type DocumentRenderer struct {
	options      RendererOptions
	engine       gopdf.GoPdf
	currentState rendererState
}

func NewDocumentRenderer(options RendererOptions) *DocumentRenderer {
	renderer := &DocumentRenderer{}
	renderer.options = options
	renderer.engine.Start(gopdf.Config{
		Unit:     gopdf.UnitPT,
		PageSize: *options.PageSize.ToRect(),
	})

	renderer.AddFontFamily(standardFontFamily)
	renderer.SetFont(standardFont)

	renderer.currentState.Font = standardFont
	renderer.currentState.Offset = renderer.GetCurrentPosition()

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

func (r *DocumentRenderer) SetX(value float64) {
	r.engine.SetX(value)
}

func (r *DocumentRenderer) SetY(value float64) {
	r.engine.SetY(value)
}

func (r *DocumentRenderer) AddX(value float64) {
	r.engine.SetX(r.engine.GetX() + value)
}

func (r *DocumentRenderer) AddY(value float64) {
	r.engine.SetY(r.engine.GetY() + value)
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
		r.setFont(*style.Font, true)
		defer r.SetFont(r.currentState.Font)
	}

	//TODO: Implement dynamic boundries based on text size when boundries are not set or set to 0
	if len(style.Borders) > 0 {
		r.BoxWithBorders(*style.Boundries, style.Borders...)
	}

	var texts []string
	var err error
	if len(text) > 0 {
		texts, err = r.SplitText(text, style)
		if err != nil {
			return err
		}
	}

	position := r.GetCurrentPosition()
	defer r.engine.SetNewXY(
		position.Y,
		position.X+style.Boundries.Width+r.options.DefaultSeparatorSize,
		0,
	)

	r.SetY(r.engine.GetY() + style.Padding.Top)

	if !style.Multiline {
		text := texts[0]

		if len(texts) > 1 && len(style.Overflow) > 0 {
			text = text[:len(text)-len(style.Overflow)] + style.Overflow
		}

		r.engine.SetX(position.X + style.Padding.Left)
		return r.engine.CellWithOption(
			nil,
			text,
			gopdf.CellOption{
				Align: int(style.Alignment),
			},
		)
	}

	textHeight, _ := r.engine.MeasureCellHeightByText(text)
	for index, text := range texts {
		if index > 0 {
			r.engine.Br(textHeight)
		}

		r.engine.SetX(position.X + style.Padding.Left)
		r.engine.CellWithOption(nil, text, gopdf.CellOption{
			Align: int(style.Alignment),
		})
	}

	return nil
}

func (r *DocumentRenderer) SplitText(text string, style *TextStyle) ([]string, error) {
	var texts []string
	var err error

	boundries := style.Boundries.WithPadding(style.Padding)

	if style.WordWrap {
		texts, err = r.engine.SplitTextWithWordWrap(text, boundries.Width)
	} else {
		texts, err = r.engine.SplitText(text, boundries.Width)
	}

	if len(texts) == 0 {
		return texts, err
	}

	textHeight, _ := r.engine.MeasureCellHeightByText(texts[0])
	limit := int(math.Trunc(boundries.Height / textHeight))

	if limit < len(texts) {
		texts = texts[:limit]
		if len(style.Overflow) > 0 {
			textsLength := len(texts) - 1

			lastText := texts[textsLength]
			lastTextLength := len(lastText)
			lastText = lastText[:lastTextLength-len(style.Overflow)] + style.Overflow

			texts[textsLength] = lastText
		}
	}

	return texts, err
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
		size.Width = r.engine.GetX()
	}

	if size.Height == 0 {
		size.Height = r.engine.GetY()
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

func (r *DocumentRenderer) Box(size Size, lineOptions *LineOptions) {
	position := r.GetCurrentPosition()

	if lineOptions == nil {
		lineOptions = &LineOptions{
			Style:       LineStyleSolid,
			StrokeWidth: 0,
			Color:       &Color{0, 0, 0},
		}
	}

	r.VerticalLine(size.Height+position.Y, lineOptions)

	offset := NewOffset(size.Width+position.X, position.Y)
	r.VerticalLineWithOffset(size.Height+position.Y, offset, lineOptions)

	r.HorizontalLine(size.Width+position.X, lineOptions)

	offset = NewOffset(position.X, size.Height+position.Y)
	r.HorizontalLineWithOffset(size.Width+position.X, offset, lineOptions)
}

func (r *DocumentRenderer) BoxWithBorders(size Size, borders ...Border) {
	position := r.GetCurrentPosition()
	for _, border := range borders {
		if border.Side == 0 {
			continue
		}

		if border.Side&Left != 0 {
			r.VerticalLine(size.Height+position.Y, &border.Options)
		}

		if border.Side&Right != 0 {
			offset := NewOffset(size.Width+position.X, position.Y)
			r.VerticalLineWithOffset(size.Height+position.Y, offset, &border.Options)
		}

		if border.Side&Top != 0 {
			r.HorizontalLine(size.Width+position.X, &border.Options)
		}

		if border.Side&Bottom != 0 {
			offset := NewOffset(position.X, size.Height+position.Y)
			r.HorizontalLineWithOffset(size.Width+position.X, offset, &border.Options)
		}
	}
}
