package main

import "fmt"

func RenderReport() {
	renderer := NewDocumentRenderer(RendererOptions{
		PageSize:             PageSizeA4,
		DefaultSeparatorSize: 8,
	})
	renderer.StartNewDocument()
	fmt.Printf("%v", renderer.GetCurrentOffset())
	report := List{
		Axis:          VerticalAxis,
		MainAxisSize:  PageSizeA4.Height - renderer.GetY()*2,
		CrossAxisSize: PageSizeA4.Width - renderer.GetX()*2,
		Padding:       NewVerticalMargin(15, 15),
		Borders:       []Border{NewBorder()},
		Justify:       JustifyContentSpaceBetween,
		Children: []Element{
			&Text{
				Value: "SDSD",
				Style: TextStyle{
					Alignment: BottomAlignment | RightAlignment,
					WordWrap:  true,
					Boundries: &Size{Width: 400, Height: 100},
					Padding:   NewHorizontalMargin(10, 10),
					Borders:   []Border{NewBorder()},
				},
			},
			&List{
				Axis:          HorizontalAxis,
				MainAxisSize:  PageSizeA4.Height - 300,
				CrossAxisSize: PageSizeA4.Width - 400,
				Borders:       []Border{NewBorder()},
				Children: []Element{
					&Text{
						Value: "XXXX",
						Style: TextStyle{
							Boundries: &Size{Width: 100, Height: 100},
						},
					},
					&Text{
						Value: "SDSD",
						Style: TextStyle{
							Boundries: &Size{Width: 100, Height: 100},
						},
					},
				},
			},
		},
	}

	report.Render(renderer)
	err := renderer.engine.WritePdf("temp/zhreport.pdf")
	if err != nil {
		panic(err)
	}
}
