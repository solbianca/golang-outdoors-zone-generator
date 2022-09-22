package layers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zone/core/tiles"
	"zone/zone/noise"
)

type Plants struct {
	config      *LayerConfig
	noise       *noise.NoiseMap
	tiles       *tiles.Collection
	reliefSigns *LayerSigns
	signs       *LayerSigns

	img *ebiten.Image
}

func NewPlants(config *LayerConfig, noise *noise.NoiseMap, tiles *tiles.Collection) *Plants {
	layer := &Plants{config: config, noise: noise, tiles: tiles}

	layer.drawInternal()

	return layer
}

func (p *Plants) Update() error {
	return nil
}

func (p *Plants) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(p.img, op)
}

func (p *Plants) SetReliefSigns(reliefSigns *LayerSigns) {
	p.reliefSigns = reliefSigns

	p.drawInternal()
}

func (p *Plants) SetNoise(noise *noise.NoiseMap) {
	p.noise = noise

	p.drawInternal()
}

func (p *Plants) GetSigns() *LayerSigns {
	return p.signs
}

func (p *Plants) drawInternal() {
	signs := newTreeLayerSignsByNoiseMap(p.noise, p.config.columns, p.config.rows)
	img := ebiten.NewImage(p.config.width, p.config.height)

	if p.reliefSigns == nil {
		p.signs = signs
		p.img = img
		return
	}

	for column := 0; column < p.config.columns; column++ {
		for row := 0; row < p.config.rows; row++ {
			tileSigns := signs.getTileSigns(column, row)

			if tileSigns == nil {
				continue
			}

			if !p.reliefSigns.has(column, row, GroundGrassSign) {
				continue
			}

			for _, tileSign := range tileSigns {
				var sprite *ebiten.Image
				var padding = 0
				switch tileSign {
				case TreeSign:
					sprite, padding = getRandomTree()
				case FlowerSign:
					sprite = getRandomFlower()
				case GrassSign:
					sprite = getRandomGrass()
				}

				x := float64((column * p.config.tileSize))
				y := float64(row*p.config.tileSize - padding)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(x, y)
				img.DrawImage(sprite, op)
			}
		}
	}

	p.img = img
}
