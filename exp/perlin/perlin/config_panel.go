package perlin

import (
	"arena/exp/perlin/perlin/assets"
	"fmt"
	"github.com/blizzy78/ebitenui/image"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	defaultSeed      = 1000
	defaultPointSize = 16
	defaultFrequency = 400.0
	defaultAlpha     = 3.0
	defaultBeta      = 3.0
	defaultN         = 3
)

var (
	randomGenerator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

type ConfigPanel struct {
	Panel     *Panel
	resources *resources
	ui        *ebitenui.UI

	textInput *SizedTextInput

	Seed                   int64
	Frequency, Alpha, Beta float64
	PointSize, N           int
}

func NewConfigPanel(windowWidth, windowHeight int, globalX, globalY float64, width, height int) *ConfigPanel {
	panel := NewPanel(windowWidth, windowHeight, globalX, globalY, width, height)

	configPanel := &ConfigPanel{Panel: panel, ui: &ebitenui.UI{}}
	configPanel.Seed = defaultSeed
	configPanel.PointSize = defaultPointSize
	configPanel.Frequency = defaultFrequency
	configPanel.Alpha = defaultAlpha
	configPanel.Beta = defaultBeta
	configPanel.N = defaultN
	configPanel.resources = newResources()

	configPanel.init()

	return configPanel
}

func (c *ConfigPanel) init() {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: false,
					StretchVertical:   false,
				},
			),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
	)

	rootContainer.AddChild(seedInputRow(c.resources, c))
	rootContainer.AddChild(sliderRowOne(c.resources, c))
	rootContainer.AddChild(sliderRowTwo(c.resources, c))

	c.ui.Container = rootContainer
}

func (c *ConfigPanel) Draw(screen *ebiten.Image) {
	c.ui.Draw(screen)
}

func (c *ConfigPanel) Update() {
	c.ui.Update()
}

func terraformingButton(resources *resources) *widget.Button {
	treeImg, _ := assets.GetImage("ui:button:terraforming")
	w, h := treeImg.Size()
	treeImgNineSLice := image.NewNineSlice(treeImg, [3]int{w, 0, 0}, [3]int{h, 0, 0})
	treeButtonImg := &widget.ButtonImage{
		Idle:     treeImgNineSLice,
		Hover:    treeImgNineSLice,
		Pressed:  treeImgNineSLice,
		Disabled: treeImgNineSLice,
	}
	// construct a button
	button := widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),

		// specify the images to use
		widget.ButtonOpts.Image(treeButtonImg),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				println(fmt.Sprintf("tree button"))
			},
		),
	)

	return button
}

func treeButton(resources *resources) *widget.Button {
	treeImg, _ := assets.GetImage("ui:button:tree")
	w, h := treeImg.Size()
	treeImgNineSLice := image.NewNineSlice(treeImg, [3]int{w, 0, 0}, [3]int{h, 0, 0})
	treeButtonImg := &widget.ButtonImage{
		Idle:     treeImgNineSLice,
		Hover:    treeImgNineSLice,
		Pressed:  treeImgNineSLice,
		Disabled: treeImgNineSLice,
	}
	// construct a button
	button := widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),

		// specify the images to use
		widget.ButtonOpts.Image(treeButtonImg),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				println(fmt.Sprintf("tree button"))
			},
		),
	)

	return button
}

func newTextInput(resources *resources) *widget.TextInput {
	textInput := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
		widget.TextInputOpts.Placeholder("Noise seed"),
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Stretch:  true,
					MaxWidth: 300,
				},
			),
		),
		widget.TextInputOpts.Image(resources.textInput.img),
		widget.TextInputOpts.Color(resources.textInput.color),
		widget.TextInputOpts.Padding(resources.textInput.padding),
		widget.TextInputOpts.Face(resources.textInput.face),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(resources.textInput.face, 2),
		),
	)

	textInput.InputText = strconv.FormatInt(defaultSeed, 10)

	return textInput
}

func newOkButton(resources *resources, configPanel *ConfigPanel) *widget.Button {
	// construct a button
	button := widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),

		// specify the images to use
		widget.ButtonOpts.Image(resources.button.img),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text(
			"OK", resources.button.face, &widget.ButtonTextColor{
				Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
			},
		),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				seed, _ := strconv.ParseInt(configPanel.textInput.InputText, 10, 64)
				println(fmt.Sprintf("new seed [%d]", seed))
				configPanel.Seed = seed
			},
		),
	)

	return button
}

func newRandomSeedButton(resources *resources, configPanel *ConfigPanel) *widget.Button {
	// construct a button
	button := widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),

		// specify the images to use
		widget.ButtonOpts.Image(resources.button.img),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text(
			"Random seed", resources.button.face, &widget.ButtonTextColor{
				Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
			},
		),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				min := 1
				max := 999
				seed := randomGenerator.Intn(max-min+1) + min
				println(fmt.Sprintf("New random seed [%d]", seed))
				configPanel.Seed = int64(seed)

				configPanel.textInput.InputText = strconv.Itoa(seed)
			},
		),
	)

	return button
}

func newPointSizeSlider(resources *resources, configPanel *ConfigPanel) *widget.Container {
	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(10),
			),
		),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	var text *widget.Label

	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
		widget.SliderOpts.MinMax(1, 32),
		widget.SliderOpts.Images(resources.slider.trackImage, resources.slider.handleImg),
		widget.SliderOpts.HandleSize(resources.slider.handleSize),
		widget.SliderOpts.PageSizeFunc(
			func() int {
				return 3
			},
		),
		widget.SliderOpts.ChangedHandler(
			func(args *widget.SliderChangedEventArgs) {
				println(fmt.Sprintf("Slider value [%d]", args.Current))
				text.Label = fmt.Sprintf("%d", args.Current)

				configPanel.PointSize = args.Current
			},
		),
	)

	text = widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), resources.label.face, resources.label.color),
	)

	labelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text("Point Size", resources.label.face, resources.label.color),
	)

	slider.Current = defaultPointSize

	sliderContainer.AddChild(labelText)
	sliderContainer.AddChild(slider)
	sliderContainer.AddChild(text)

	return sliderContainer
}

func newFrequencySlider(resources *resources, configPanel *ConfigPanel) *widget.Container {
	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(10),
			),
		),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	var valueText *widget.Label

	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
		widget.SliderOpts.MinMax(10, 1000),
		widget.SliderOpts.Images(resources.slider.trackImage, resources.slider.handleImg),
		widget.SliderOpts.HandleSize(resources.slider.handleSize),
		widget.SliderOpts.PageSizeFunc(
			func() int {
				return 150
			},
		),
		widget.SliderOpts.ChangedHandler(
			func(args *widget.SliderChangedEventArgs) {
				println(fmt.Sprintf("frequency slider value [%d]", args.Current))
				valueText.Label = fmt.Sprintf("%d", args.Current)

				configPanel.Frequency = float64(args.Current)
			},
		),
	)

	valueText = widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), resources.label.face, resources.label.color),
	)

	slider.Current = defaultFrequency

	labelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text("frequency", resources.label.face, resources.label.color),
	)

	sliderContainer.AddChild(labelText)
	sliderContainer.AddChild(slider)
	sliderContainer.AddChild(valueText)

	return sliderContainer
}

func newAlphaSlider(resources *resources, configPanel *ConfigPanel) *widget.Container {
	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(10),
			),
		),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	var valueText *widget.Label

	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
		widget.SliderOpts.MinMax(1, 10),
		widget.SliderOpts.Images(resources.slider.trackImage, resources.slider.handleImg),
		widget.SliderOpts.HandleSize(resources.slider.handleSize),
		widget.SliderOpts.PageSizeFunc(
			func() int {
				return 2
			},
		),
		widget.SliderOpts.ChangedHandler(
			func(args *widget.SliderChangedEventArgs) {
				println(fmt.Sprintf("alpha slider value [%d]", args.Current))
				valueText.Label = fmt.Sprintf("%d", args.Current)

				configPanel.Alpha = float64(args.Current)
			},
		),
	)

	valueText = widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), resources.label.face, resources.label.color),
	)

	slider.Current = defaultAlpha

	labelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text("alpha", resources.label.face, resources.label.color),
	)

	sliderContainer.AddChild(labelText)
	sliderContainer.AddChild(slider)
	sliderContainer.AddChild(valueText)

	return sliderContainer
}

func newBetaSlider(resources *resources, configPanel *ConfigPanel) *widget.Container {
	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(10),
			),
		),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	var valueText *widget.Label

	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
		widget.SliderOpts.MinMax(1, 10),
		widget.SliderOpts.Images(resources.slider.trackImage, resources.slider.handleImg),
		widget.SliderOpts.HandleSize(resources.slider.handleSize),
		widget.SliderOpts.PageSizeFunc(
			func() int {
				return 2
			},
		),
		widget.SliderOpts.ChangedHandler(
			func(args *widget.SliderChangedEventArgs) {
				println(fmt.Sprintf("beta slider value [%d]", args.Current))
				valueText.Label = fmt.Sprintf("%d", args.Current)

				configPanel.Beta = float64(args.Current)
			},
		),
	)

	valueText = widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), resources.label.face, resources.label.color),
	)

	slider.Current = defaultBeta

	labelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text("beta", resources.label.face, resources.label.color),
	)

	sliderContainer.AddChild(labelText)
	sliderContainer.AddChild(slider)
	sliderContainer.AddChild(valueText)

	return sliderContainer
}

func newNSlider(resources *resources, configPanel *ConfigPanel) *widget.Container {
	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(10),
			),
		),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	var valueText *widget.Label

	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
		widget.SliderOpts.MinMax(1, 10),
		widget.SliderOpts.Images(resources.slider.trackImage, resources.slider.handleImg),
		widget.SliderOpts.HandleSize(resources.slider.handleSize),
		widget.SliderOpts.PageSizeFunc(
			func() int {
				return 2
			},
		),
		widget.SliderOpts.ChangedHandler(
			func(args *widget.SliderChangedEventArgs) {
				println(fmt.Sprintf("n slider value [%d]", args.Current))
				valueText.Label = fmt.Sprintf("%d", args.Current)

				configPanel.N = args.Current
			},
		),
	)

	valueText = widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), resources.label.face, resources.label.color),
	)

	slider.Current = defaultN

	labelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text("n", resources.label.face, resources.label.color),
	)

	sliderContainer.AddChild(labelText)
	sliderContainer.AddChild(slider)
	sliderContainer.AddChild(valueText)

	return sliderContainer
}

//func openWindow(resources *resources, configPanel *ConfigPanel, coreui *ebitenui.UI) {
//	var rw ebitenui.RemoveWindowFunc
//
//	modalContainer := widget.NewContainer(
//		widget.ContainerOpts.BackgroundImage(resources.panel.img),
//		widget.ContainerOpts.Layout(
//			widget.NewRowLayout(
//				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
//				widget.RowLayoutOpts.Padding(resources.panel.padding),
//				widget.RowLayoutOpts.Spacing(15),
//			),
//		),
//	)
//
//	modalContainer.AddChild(seedInputRow(resources, configPanel))
//	modalContainer.AddChild(sliderRowOne(resources, configPanel))
//	modalContainer.AddChild(sliderRowTwo(resources, configPanel))
//
//	bc := widget.NewContainer(
//		widget.ContainerOpts.Layout(
//			widget.NewRowLayout(
//				widget.RowLayoutOpts.Spacing(15),
//			),
//		),
//	)
//	modalContainer.AddChild(bc)
//
//	cb := widget.NewButton(
//		widget.ButtonOpts.Image(resources.button.img),
//		widget.ButtonOpts.TextPadding(resources.button.padding),
//		widget.ButtonOpts.Text("Close", resources.button.face, resources.button.color),
//		widget.ButtonOpts.ClickedHandler(
//			func(args *widget.ButtonClickedEventArgs) {
//				rw()
//			},
//		),
//	)
//	bc.AddChild(cb)
//
//	w := widget.NewWindow(
//		widget.WindowOpts.Modal(),
//		widget.WindowOpts.Contents(modalContainer),
//	)
//
//	ww, wh := ebiten.WindowSize()
//	r := image.Rect(0, 0, ww*3/4, wh/3)
//	r = r.Add(image.Point{ww / 4 / 2, wh * 2 / 3 / 2})
//	w.SetLocation(r)
//
//	rw = coreui.AddWindow(w)
//}

func seedInputRow(resources *resources, configPanel *ConfigPanel) *widget.Container {
	seedTextInput := NewSizedTextInput(newTextInput(resources), 150, 0)
	okButton := newOkButton(resources, configPanel)
	randomSeedButton := newRandomSeedButton(resources, configPanel)

	seedContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(15),
				widget.RowLayoutOpts.Padding(resources.panel.padding),
			),
		),
	)

	seedLabelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text("seed", resources.label.face, resources.label.color),
	)

	seedContainer.AddChild(seedLabelText)
	seedContainer.AddChild(seedTextInput)
	seedContainer.AddChild(okButton)
	seedContainer.AddChild(randomSeedButton)

	return seedContainer
}

func sliderRowOne(resources *resources, configPanel *ConfigPanel) *widget.Container {
	pointSizeSlider := newPointSizeSlider(resources, configPanel)
	frequencySlider := newFrequencySlider(resources, configPanel)

	slidersContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(15),
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
	)

	slidersContainer.AddChild(pointSizeSlider)
	slidersContainer.AddChild(frequencySlider)

	return slidersContainer
}

func sliderRowTwo(resources *resources, configPanel *ConfigPanel) *widget.Container {
	alpha := newAlphaSlider(resources, configPanel)
	beta := newBetaSlider(resources, configPanel)
	n := newNSlider(resources, configPanel)

	slidersContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(15),
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
	)

	slidersContainer.AddChild(alpha)
	slidersContainer.AddChild(beta)
	slidersContainer.AddChild(n)

	return slidersContainer
}
