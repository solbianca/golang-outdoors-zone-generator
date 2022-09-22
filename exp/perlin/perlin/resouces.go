package perlin

import (
	"arena/exp/perlin/perlin/assets"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

const (
	backgroundColor = "131a22"

	textIdleColor     = "dff4ff"
	textDisabledColor = "5a7a91"

	labelIdleColor     = textIdleColor
	labelDisabledColor = textDisabledColor

	buttonIdleColor     = textIdleColor
	buttonDisabledColor = labelDisabledColor

	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"

	headerColor = textIdleColor

	textInputCaretColor         = "e7c34b"
	textInputDisabledCaretColor = "766326"

	toolTipColor = backgroundColor

	separatorColor = listDisabledSelectedBackground
)

type resources struct {
	panel     *panel
	button    *button
	textInput *textInput
	slider    *slider
	label     *label
	text      *text
}

func newResources() *resources {
	fontFace, _ := assets.GetFont("mana", 18)
	idle, _ := assets.GetImage("liteui:input")
	btn, _ := assets.GetImage("liteui:btn:dark")
	sliderHandle, _ := assets.GetImage("liteui:btn:radio")
	barFrame, _ := assets.GetImage("liteui:bar-frame")
	panelImg, _ := assets.GetImage("liteui:panel")

	return &resources{
		panel: &panel{
			img: image.NewNineSlice(panelImg, [3]int{8, 48, 8}, [3]int{8, 48, 8}),
			padding: widget.Insets{
				Top:    5,
				Bottom: 5,
				Left:   5,
				Right:  5,
			},
		},
		button: &button{
			img: &widget.ButtonImage{
				Idle: image.NewNineSlice(btn, [3]int{5, 38, 5}, [3]int{5, 6, 5}),
			},
			face: fontFace.Face,
			padding: widget.Insets{
				Left:   10,
				Right:  10,
				Top:    10,
				Bottom: 6,
			},
			color: &widget.ButtonTextColor{
				Idle:     hexToColor(textIdleColor),
				Disabled: hexToColor(textDisabledColor),
			},
		},
		textInput: &textInput{
			img: &widget.TextInputImage{
				Idle:     image.NewNineSlice(idle, [3]int{5, 76, 5}, [3]int{5, 4, 5}),
				Disabled: image.NewNineSlice(idle, [3]int{5, 76, 5}, [3]int{5, 4, 5}),
			},
			color: &widget.TextInputColor{
				Idle:          hexToColor(textIdleColor),
				Disabled:      hexToColor(textDisabledColor),
				Caret:         hexToColor(textInputCaretColor),
				DisabledCaret: hexToColor(textInputDisabledCaretColor),
			},
			face: fontFace.Face,
			padding: widget.Insets{
				Top:    10,
				Bottom: 0,
				Left:   10,
				Right:  10,
			},
		},
		slider: &slider{
			trackImage: &widget.SliderTrackImage{
				Idle:  image.NewNineSlice(barFrame, [3]int{2, 43, 2}, [3]int{2, 1, 2}),
				Hover: image.NewNineSlice(barFrame, [3]int{2, 43, 2}, [3]int{2, 1, 2}),
			},
			handleImg: &widget.ButtonImage{
				Idle:     image.NewNineSlice(sliderHandle, [3]int{3, 2, 3}, [3]int{3, 2, 3}),
				Hover:    image.NewNineSlice(sliderHandle, [3]int{3, 2, 3}, [3]int{3, 2, 3}),
				Pressed:  image.NewNineSlice(sliderHandle, [3]int{3, 2, 3}, [3]int{3, 2, 3}),
				Disabled: image.NewNineSlice(sliderHandle, [3]int{3, 2, 3}, [3]int{3, 2, 3}),
			},
			handleSize: 20,
		},
		label: &label{
			face: fontFace.Face,
			color: &widget.LabelColor{
				Idle:     hexToColor(labelIdleColor),
				Disabled: hexToColor(labelDisabledColor),
			},
		},
		text: &text{
			face:  fontFace.Face,
			color: hexToColor(labelIdleColor),
		},
	}
}

type panel struct {
	img     *image.NineSlice
	padding widget.Insets
}

type button struct {
	img     *widget.ButtonImage
	face    font.Face
	padding widget.Insets
	color   *widget.ButtonTextColor
}

type textInput struct {
	img     *widget.TextInputImage
	color   *widget.TextInputColor
	face    font.Face
	padding widget.Insets
}

type slider struct {
	trackImage *widget.SliderTrackImage
	handleImg  *widget.ButtonImage
	handleSize int
}

type label struct {
	face  font.Face
	color *widget.LabelColor
}
type text struct {
	face  font.Face
	color color.Color
}

func hexToColor(h string) color.Color {
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}
