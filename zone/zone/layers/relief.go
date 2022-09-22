package layers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zone/core/resources"
	"zone/core/tiles"
	"zone/zone/noise"
)

var (
	rocks = map[sign]bool{
		RocksSmallSign: true,
		RocksHugeSign:  true,
	}
)

type Relief struct {
	config *LayerConfig
	noise  *noise.NoiseMap
	tiles  *tiles.Collection
	signs  *LayerSigns

	img *ebiten.Image
}

func NewRelief(config *LayerConfig, noise *noise.NoiseMap, tiles *tiles.Collection) *Relief {
	layer := &Relief{config: config, noise: noise, tiles: tiles, signs: newLayerSigns()}

	layer.drawInternal()

	return layer
}

func (r *Relief) Update() error {
	return nil
}

func (r *Relief) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(r.img, op)
}

func (r *Relief) SetNoise(noise *noise.NoiseMap) {
	r.noise = noise

	r.drawInternal()
}

func (r *Relief) GetSigns() *LayerSigns {
	return r.signs
}

func (r *Relief) drawInternal() {
	signs := newReliefLayerSignsByNoiseMap(r.noise, r.config.columns, r.config.rows)
	img := ebiten.NewImage(r.config.width, r.config.height)

	for column := 0; column < r.config.columns; column++ {
		for row := 0; row < r.config.rows; row++ {
			tileSigns := signs.getTileSigns(column, row)

			if tileSigns == nil {
				continue
			}

			for _, tileSign := range tileSigns {
				drawGround(img, tileSign, column, row, r.config.tileSize)
				drawWater(img, signs, tileSign, column, row, r.config.tileSize)
				drawRocks(img, tileSign, column, row, r.config.tileSize)
				drawMountains(img, tileSign, column, row, r.config.tileSize)
			}
		}
	}

	r.signs = signs
	r.img = img
}

func drawGround(screen *ebiten.Image, sign sign, column, row, tileSize int) {
	if sign != GroundGrassSign {
		return
	}

	x := float64(column * tileSize)
	y := float64(row * tileSize)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(getRandomGroundSprite(), op)
}

func drawWater(screen *ebiten.Image, signs *LayerSigns, sign sign, column, row, tileSize int) {
	if sign != WaterSymbol {
		return
	}

	sprite := getRandomWaterSprite()

	x := float64(column * tileSize)
	y := float64(row * tileSize)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(sprite, op)
}

func drawRocks(screen *ebiten.Image, sign sign, column, row, tileSize int) {
	if _, ok := rocks[sign]; !ok {
		return
	}

	var sprite *ebiten.Image
	if sign == RocksSmallSign {
		sprite = getRandomSandstoneSmall()
	} else if sign == RocksHugeSign {
		sprite = getRandomSandstoneHuge()
	}

	x := float64(column * tileSize)
	y := float64(row * tileSize)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(sprite, op)
}

func drawMountains(screen *ebiten.Image, sign sign, column, row, tileSize int) {
	if sign != MountainSign {
		return
	}

	x := float64(column * tileSize)
	y := float64(row * tileSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(resources.GetSpriteOrPanic("ground:grass:1").Original(), op)

	mountainsSpritesheet := getMountainSpritesheet()
	sprite := mountainsSpritesheet.Sprite(1, 6).Original()
	screen.DrawImage(sprite, op)
}
