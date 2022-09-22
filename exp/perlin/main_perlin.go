package main

import (
	"arena/exp/perlin/perlin"
	"arena/exp/perlin/perlin/assets"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	windowWidth  = 1600
	windowHeight = 900

	mainPanel   *perlin.MainPanel
	configPanel *perlin.ConfigPanel

	// debugRenderer *renders.DebugRenderer
)

type game struct{}

func (g game) Update() error {
	configPanel.Update()

	if configPanel.Seed != mainPanel.GetSeed() {
		mainPanel.SetSeed(configPanel.Seed)
	}
	if configPanel.PointSize != mainPanel.GetPointSize() {
		mainPanel.SetPointSize(configPanel.PointSize)
	}
	if configPanel.Frequency != mainPanel.GetFrequency() {
		mainPanel.SetFrequency(configPanel.Frequency)
	}
	if configPanel.Alpha != mainPanel.GetAlpha() {
		mainPanel.SetAlpha(configPanel.Alpha)
	}
	if configPanel.Beta != mainPanel.GetBeta() {
		mainPanel.SetBeta(configPanel.Beta)
	}
	if configPanel.N != mainPanel.GetN() {
		mainPanel.SetN(configPanel.N)
	}

	mainPanel.Update()

	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	mainPanel.Draw(screen)
	configPanel.Draw(screen)

	// debugRenderer.Process(screen)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Perlin Noise Demo")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(true)

	loadAssets()
	// debugRenderer = renders.NewDebugRenderer()

	configPanel = perlin.NewConfigPanel(windowWidth, windowHeight, 0.0, 0.0, 1600, 120)
	mainPanel = perlin.NewMainPanel(windowWidth, windowHeight, 0.0, 120.0, 1600, 900)

	game := game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func loadAssets() {
	assets.LoadFont("mana", "exp/perlin/resources/fonts/manaspace.ttf")
	assets.LoadFont("pixel", "exp/perlin/resources/fonts/pixel-nes.otf")

	assets.LoadImage("liteui:input", "exp/perlin/resources/liteui/Input Field 1.png")
	assets.LoadImage("liteui:btn:dark", "exp/perlin/resources/liteui/Big Dark Btn.png")
	assets.LoadImage("liteui:btn:circular", "exp/perlin/resources/liteui/Circular Btn.png")
	assets.LoadImage("liteui:btn:radio", "exp/perlin/resources/liteui/Radio Btn Frame.png")
	assets.LoadImage("liteui:bar-frame", "exp/perlin/resources/liteui/Bar Frame.png")

	assets.LoadImage("liteui:panel", "exp/perlin/resources/liteui/panel.png")

	assets.LoadImage("ui:button:terraforming", "exp/perlin/resources/terraforming.png")
	assets.LoadImage("ui:button:tree", "exp/perlin/resources/tree.png")
}
