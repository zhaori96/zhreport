package main

import (
	"github.com/signintech/gopdf"
)

var (
	standardFontFamily = FontFamily{
		Name:       "roboto",
		Path:       "assets/fonts/roboto",
		Regular:    "Roboto-Regular.ttf",
		Italic:     "Roboto-Italic.ttf",
		Bold:       "Roboto-Bold.ttf",
		BoldItalic: "Roboto-BoldItalic.ttf",
	}

	standardFont = Font{
		Family: "roboto",
		Size:   8,
	}
)

type FontFamily struct {
	Name       string
	Path       string
	Regular    string
	Italic     string
	Bold       string
	BoldItalic string
}

func (f *FontFamily) HasRegularStyle() bool {
	return len(f.Regular) > 0
}

func (f *FontFamily) HasItalicStyle() bool {
	return len(f.Italic) > 0
}

func (f *FontFamily) HasBoldStyle() bool {
	return len(f.Bold) > 0
}

func (f *FontFamily) HasBoldItalicStyle() bool {
	return len(f.BoldItalic) > 0
}

func FontFamilyFromPath(name string, path string) (*FontFamily, error) {
	//TODO: Create some way to load a font family from a folder.
	return nil, nil
}

type Font struct {
	Family string
	Size   float64
	Style  *FontStyle
}

type FontStyle struct {
	Bold      bool
	Italic    bool
	Underline bool
}

func NewFontStyle(bold, italic, underline bool) FontStyle {
	return FontStyle{
		Bold:      bold,
		Italic:    italic,
		Underline: underline,
	}
}

func (f *FontStyle) Combine() int {
	if f == nil {
		return 0
	}

	var combination int
	if f.Bold {
		combination = combination | gopdf.Bold
	}

	if f.Italic {
		combination = combination | gopdf.Italic
	}

	if f.Underline {
		combination = combination | gopdf.Underline
	}

	return combination
}
