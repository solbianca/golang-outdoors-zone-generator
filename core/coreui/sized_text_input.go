package coreui

import "github.com/blizzy78/ebitenui/widget"

type SizedTextInput struct {
	*widget.TextInput

	width  int
	height int
}

func NewSizedTextInput(textInput *widget.TextInput, width int, height int) *SizedTextInput {
	return &SizedTextInput{TextInput: textInput, width: width, height: height}
}

func (t *SizedTextInput) PreferredSize() (int, int) {
	targetWidth, targetHeight := t.TextInput.PreferredSize()

	if t.width != 0 {
		targetWidth = t.width
	}

	if t.height != 0 {
		targetHeight = t.height
	}

	return targetWidth, targetHeight
}
