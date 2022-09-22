package zone

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"zone/zone/ui"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 768
	TileSize     = 16
)

type GameObject interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Game struct {
	configPanel *ui.ConfigPanel
	regionPanel *ui.ZonePanel
}

func NewGame() *Game {
	columns, rows := 1024/TileSize, 640/TileSize
	observer := ui.NewConfigObserver(columns, rows)

	g := &Game{
		configPanel: ui.NewConfigPanel(984, 120, observer),
		regionPanel: ui.NewZonePanel(0, 140, 1024, 640, columns, rows, TileSize),
	}
	observer.Config = g.configPanel
	observer.Zone = g.regionPanel

	return g
}

func (g *Game) Update() error {
	_ = g.configPanel.Update()
	_ = g.regionPanel.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.configPanel.Draw(screen)
	g.regionPanel.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f, FPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
