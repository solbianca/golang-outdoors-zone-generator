package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zone/zone/zone"
)

type ZonePanel struct {
	x, y          float64
	width, height int
	columns, rows int
	tileSize      int

	zone *zone.Zone
}

func NewZonePanel(x, y float64, width, height, columns, rows, tileSize int) *ZonePanel {
	panel := &ZonePanel{x: x, y: y, width: width, height: height, columns: columns, rows: rows, tileSize: tileSize}

	panel.zone = zone.NewZone(width, height, columns, rows, tileSize)

	return panel
}

func (z *ZonePanel) Update() error {
	err := z.zone.Update()
	if err != nil {
		return err
	}

	return nil
}

func (z *ZonePanel) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(z.width, z.height)
	z.zone.Draw(img)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(z.x, z.y)

	screen.DrawImage(img, op)
}
