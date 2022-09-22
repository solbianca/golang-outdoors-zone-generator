package ui

import (
	"fmt"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"strconv"
	"zone/core/coreui"
)

const (
	relief = "relief"
	plants = "plants"
)

const (
	defaultSeed      = 1000
	defaultPointSize = 16
	defaultFrequency = 400.0
	defaultAlpha     = 3.0
	defaultBeta      = 3.0
	defaultN         = 3
)

type layerConfig struct {
	alias string

	noiseConfig       *noiseConfig
	seedInput         *coreui.SizedTextInput
	layerConfigButton *widget.Button
	inputs            *widget.Container
}

func newLayerConfig(alias string) *layerConfig {
	return &layerConfig{
		alias:       alias,
		noiseConfig: newDefaultNoiseConfig(),
	}
}

type noiseConfig struct {
	seed                   int64
	frequency, alpha, beta float64
	pointSize, n           int
}

func newDefaultNoiseConfig() *noiseConfig {
	return &noiseConfig{
		seed:      defaultSeed,
		frequency: defaultFrequency,
		alpha:     defaultAlpha,
		beta:      defaultBeta,
		pointSize: defaultPointSize,
		n:         defaultN,
	}
}

type ConfigPanel struct {
	width, height int
	resources     *gameResources
	ui            *ebitenui.UI

	noiseConfig *noiseConfig

	layers             map[string]*layerConfig
	removeActiveInputs widget.RemoveChildFunc

	observer *ConfigObserver
}

func NewConfigPanel(width, height int, observer *ConfigObserver) *ConfigPanel {
	configPanel := &ConfigPanel{
		width:  width,
		height: height, ui: &ebitenui.UI{},
		observer: observer,
		layers: map[string]*layerConfig{
			relief: newLayerConfig(relief),
			plants: newLayerConfig(plants),
		},
	}

	configPanel.noiseConfig = newDefaultNoiseConfig()

	configPanel.resources = newResources()
	configPanel.init()

	return configPanel
}

func (c *ConfigPanel) init() {
	panel := coreui.NewSizedPanel(
		c.width,
		c.height,
		widget.ContainerOpts.BackgroundImage(c.resources.panel.img),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{},
			),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(
					widget.Insets{
						Top:    5,
						Bottom: 10,
						Left:   20,
						Right:  20,
					},
				),
			),
		),
	)

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
				},
			),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(
					widget.Insets{
						Top:    10,
						Bottom: 10,
						Left:   20,
						Right:  20,
					},
				),
			),
		),
	)

	panel.Container().AddChild(switchConfigPanelButtons(c.resources, c, panel))

	c.layers[relief].inputs = configInputs(c.resources, c.layers[relief], c.observer)
	c.layers[plants].inputs = configInputs(c.resources, c.layers[plants], c.observer)

	c.removeActiveInputs = panel.Container().AddChild(c.layers[relief].inputs)

	rootContainer.AddChild(panel)

	c.ui.Container = rootContainer
}

func (c *ConfigPanel) Draw(screen *ebiten.Image) {
	c.ui.Draw(screen)
}

func (c *ConfigPanel) Update() error {
	c.ui.Update()

	return nil
}

func switchConfigPanelButtons(resources *gameResources, configPanel *ConfigPanel, panel *coreui.SizedPanel) *widget.Container {
	buttonsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(
					widget.Insets{
						Top:    10,
						Bottom: 10,
						Left:   20,
						Right:  20,
					},
				),
			),
		),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	buttonsContainer.AddChild(reliefButton(resources, configPanel, panel))
	buttonsContainer.AddChild(plantsButton(resources, configPanel, panel))

	return buttonsContainer
}

func configInputs(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
	configInputsContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
				},
			),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(
					widget.Insets{
						Top:    5,
						Bottom: 10,
						Left:   20,
						Right:  20,
					},
				),
			),
		),
	)

	configInputsContainer.AddChild(seedInputRow(resources, config, observer))
	configInputsContainer.AddChild(sliderRowOne(resources, config, observer))
	configInputsContainer.AddChild(sliderRowTwo(resources, config, observer))

	return configInputsContainer
}

func reliefButton(resources *gameResources, configPanel *ConfigPanel, panel *coreui.SizedPanel) *widget.Button {
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
		//widget.ButtonOpts.Image(resources.terraformingButtonImg),
		widget.ButtonOpts.Image(resources.button.imgActive),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text(
			"RELIEF", resources.button.face, resources.button.color,
		),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				configPanel.layers[relief].layerConfigButton.Image = resources.button.imgActive
				configPanel.layers[plants].layerConfigButton.Image = resources.button.img

				configPanel.removeActiveInputs()
				configPanel.removeActiveInputs = panel.Container().AddChild(configPanel.layers[relief].inputs)

				println(fmt.Sprintf("relief button"))
			},
		),
	)

	configPanel.layers[relief].layerConfigButton = button

	return button
}

func plantsButton(resources *gameResources, configPanel *ConfigPanel, panel *coreui.SizedPanel) *widget.Button {
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
		//widget.ButtonOpts.Image(resources.treeButtonImg),

		// specify the images to use
		//widget.ButtonOpts.Image(resources.terraformingButtonImg),
		widget.ButtonOpts.Image(resources.button.img),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text(
			"PLANTS", resources.button.face, resources.button.color,
		),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				configPanel.layers[relief].layerConfigButton.Image = resources.button.img
				configPanel.layers[plants].layerConfigButton.Image = resources.button.imgActive

				configPanel.removeActiveInputs()
				configPanel.removeActiveInputs = panel.Container().AddChild(configPanel.layers[plants].inputs)

				println(fmt.Sprintf("tree button"))
			},
		),
	)

	configPanel.layers[plants].layerConfigButton = button

	return button
}

func newTextInput(resources *gameResources) *widget.TextInput {
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

func newOkButton(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Button {
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
			"OK", resources.button.face, resources.button.color,
		),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(resources.button.padding),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				seed, _ := strconv.ParseInt(config.seedInput.InputText, 10, 64)
				println(fmt.Sprintf("new seed [%d]", seed))
				config.noiseConfig.seed = seed

				observer.Notify(config.alias)
			},
		),
	)

	return button
}

func newRandomSeedButton(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Button {
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
			"Random seed", resources.button.face, resources.button.color,
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
				config.noiseConfig.seed = int64(seed)

				observer.Notify(config.alias)
				//configPanel.observer.Notify()

				config.seedInput.InputText = strconv.Itoa(seed)
			},
		),
	)

	return button
}

func newFrequencySlider(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
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

				config.noiseConfig.frequency = float64(args.Current)
				observer.Notify(config.alias)
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

func newAlphaSlider(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
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

				config.noiseConfig.alpha = float64(args.Current)
				observer.Notify(config.alias)
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

func newBetaSlider(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
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

				config.noiseConfig.beta = float64(args.Current)
				observer.Notify(config.alias)
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

func newNSlider(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
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
				config.noiseConfig.n = args.Current

				observer.Notify(config.alias)
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

func seedInputRow(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
	seedInput := coreui.NewSizedTextInput(newTextInput(resources), 150, 0)
	okButton := newOkButton(resources, config, observer)
	randomSeedButton := newRandomSeedButton(resources, config, observer)

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
	seedContainer.AddChild(seedInput)
	seedContainer.AddChild(okButton)
	seedContainer.AddChild(randomSeedButton)

	config.seedInput = seedInput

	return seedContainer
}

func sliderRowOne(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
	frequencySlider := newFrequencySlider(resources, config, observer)
	n := newNSlider(resources, config, observer)

	slidersContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(30),
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
	)

	slidersContainer.AddChild(frequencySlider)
	slidersContainer.AddChild(n)

	return slidersContainer
}

func sliderRowTwo(resources *gameResources, config *layerConfig, observer *ConfigObserver) *widget.Container {
	alpha := newAlphaSlider(resources, config, observer)
	beta := newBetaSlider(resources, config, observer)

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

	return slidersContainer
}
