package ui

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
	"image/color"
	"zone/core/resources"
)

const (
	backgroundColor = "131a22"

	textIdleColor     = "333333"
	textDisabledColor = "5a7a91"

	labelIdleColor     = textIdleColor
	labelDisabledColor = textDisabledColor

	buttonIdleColor     = "333333"
	buttonDisabledColor = "999999"

	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"

	headerColor = textIdleColor

	textInputCaretColor         = "e7c34b"
	textInputDisabledCaretColor = "766326"

	toolTipColor = backgroundColor

	separatorColor = listDisabledSelectedBackground
)

var (
	res *gameResources
)

type gameResources struct {
	panel     *panel
	button    *button
	textInput *textInput
	slider    *slider
	label     *label
	text      *text
}

func newResources() *gameResources {
	fontFace := resources.GetFontOrPanic("mana", 18)
	idle := resources.GetSpriteOrPanic("liteui:input").Original()
	sliderHandle := resources.GetSpriteOrPanic("liteui:btn:radio").Original()
	barFrame := resources.GetSpriteOrPanic("liteui:bar-frame").Original()
	panelImg := resources.GetSpriteOrPanic("liteui:panel").Original()

	return &gameResources{
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
				Idle: image.NewNineSlice(
					resources.GetSpriteOrPanic("ui:button:idle").Original(),
					[3]int{3, 10, 3},
					[3]int{3, 10, 3},
				),
				Hover: image.NewNineSlice(
					resources.GetSpriteOrPanic("ui:button:hover").Original(),
					[3]int{3, 10, 3},
					[3]int{3, 10, 3},
				),
				Pressed: image.NewNineSlice(
					resources.GetSpriteOrPanic("ui:button:pressed").Original(),
					[3]int{3, 10, 3},
					[3]int{3, 10, 3},
				),
			},
			imgActive: &widget.ButtonImage{
				Idle: image.NewNineSlice(
					resources.GetSpriteOrPanic("ui:button:active").Original(),
					[3]int{3, 10, 3},
					[3]int{3, 10, 3},
				),
			},
			face: fontFace,
			padding: widget.Insets{
				Left:   10,
				Right:  10,
				Top:    10,
				Bottom: 6,
			},
			color: &widget.ButtonTextColor{
				Idle:     hexToColor(buttonIdleColor),
				Disabled: hexToColor(buttonDisabledColor),
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
			face: fontFace,
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
			face: fontFace,
			color: &widget.LabelColor{
				Idle:     hexToColor(labelIdleColor),
				Disabled: hexToColor(labelDisabledColor),
			},
		},
		text: &text{
			face:  fontFace,
			color: hexToColor(labelIdleColor),
		},
	}
}

type panel struct {
	img     *image.NineSlice
	padding widget.Insets
}

type button struct {
	img       *widget.ButtonImage
	imgActive *widget.ButtonImage
	face      font.Face
	padding   widget.Insets
	color     *widget.ButtonTextColor
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
