package layers

import (
	"fmt"
	"zone/core/utils/random"
	"zone/zone/noise"
)

const (
	GroundGrassSign = "ground:grass"
	RocksSmallSign  = "ground:rocks:small"
	RocksHugeSign   = "ground:rocks:huge"
	WaterSymbol     = "water"
	MountainSign    = "mountain"

	GrassSign  = "grass"
	FlowerSign = "flower"
	TreeSign   = "tree"
)

type sign string

type LayerSigns struct {
	tiles           map[string][]sign
	uniqueTileSigns map[string]string
}

func newLayerSigns() *LayerSigns {
	return &LayerSigns{tiles: map[string][]sign{}, uniqueTileSigns: map[string]string{}}
}

func newReliefLayerSignsByNoiseMap(noise *noise.NoiseMap, columns, rows int) *LayerSigns {
	layerSigns := newLayerSigns()

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			point := noise.Get(column, row)
			switch {
			case point >= -1 && point < -0.5:
				layerSigns.set(column, row, WaterSymbol)
			case point >= 0.20 && point < 0.30:
				layerSigns.set(column, row, RocksSmallSign)
			case point >= 0.30 && point < 0.33:
				layerSigns.set(column, row, RocksHugeSign)
			case point >= 0.33 && point <= 1:
				layerSigns.set(column, row, MountainSign)
			default:
				layerSigns.set(column, row, GroundGrassSign)
			}
		}
	}

	return layerSigns
}

func newTreeLayerSignsByNoiseMap(noise *noise.NoiseMap, columns, rows int) *LayerSigns {
	layerSigns := newLayerSigns()

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			point := noise.Get(column, row)

			switch {
			case point >= -1 && point < 0.3:
				if random.Between(0, 100) > 95 {
					layerSigns.set(column, row, TreeSign)
				}
			case point >= 0.30 && point < 0.5:
				if random.Between(0, 100) > 65 {
					layerSigns.set(column, row, TreeSign)
				}
			case point >= 0.5 && point <= 1:
				if random.Between(0, 100) > 50 {
					layerSigns.set(column, row, TreeSign)
				}
			default:
			}

			if layerSigns.has(column, row, TreeSign) {
				continue
			}

			switch {
			case point >= -1 && point < -0.5:
				if random.Between(0, 100) > 65 {
					layerSigns.set(column, row, FlowerSign)
				}
			case point >= -0.50 && point < -0.33:
				if random.Between(0, 100) > 75 {
					layerSigns.set(column, row, FlowerSign)
				}
			case point >= -0.33 && point < 0:
				if random.Between(0, 100) > 85 {
					layerSigns.set(column, row, FlowerSign)
				}
			case point >= 0.0 && point <= 0.5:
				if random.Between(0, 100) > 95 {
					layerSigns.set(column, row, FlowerSign)
				}
			case point >= 0.5 && point <= 1:
				if random.Between(0, 100) > 100 {
					layerSigns.set(column, row, FlowerSign)
				}
			default:
			}

			if layerSigns.has(column, row, TreeSign) || layerSigns.has(column, row, FlowerSign) {
				continue
			}

			switch {
			case point >= -1 && point < -0.5:
				if random.Between(0, 100) > 95 {
					layerSigns.set(column, row, GrassSign)
				}
			case point >= -0.50 && point < -0.33:
				if random.Between(0, 100) > 85 {
					layerSigns.set(column, row, GrassSign)
				}
			case point >= -0.33 && point < 0.33:
				if random.Between(0, 100) > 65 {
					layerSigns.set(column, row, GrassSign)
				}
			case point >= 0.0 && point <= 0.33:
				if random.Between(0, 100) > 85 {
					layerSigns.set(column, row, GrassSign)
				}
			case point >= 0.5 && point <= 1:
				if random.Between(0, 100) > 95 {
					layerSigns.set(column, row, GrassSign)
				}
			default:
			}
		}
	}

	return layerSigns
}

func (l *LayerSigns) getTileSigns(column, row int) []sign {
	if signs, ok := l.tiles[l.tileAddressKey(column, row)]; ok {
		return signs
	}

	return nil
}

func (l *LayerSigns) has(column, row int, sign sign) bool {
	_, ok := l.uniqueTileSigns[l.uniqueTileSignsKey(column, row, sign)]

	return ok
}

func (l *LayerSigns) set(column, row int, s sign) {
	uniqueKey := l.uniqueTileSignsKey(column, row, s)

	if _, ok := l.uniqueTileSigns[uniqueKey]; ok {
		return
	}

	tileAddress := l.tileAddressKey(column, row)
	tileSigns, ok := l.tiles[tileAddress]

	if !ok {
		tileSigns = []sign{}
	}

	tileSigns = append(tileSigns, s)

	l.uniqueTileSigns[uniqueKey] = uniqueKey
	l.tiles[tileAddress] = tileSigns
}

func (l *LayerSigns) tileAddressKey(column, row int) string {
	return fmt.Sprintf("%d:%d", column, row)
}

func (l *LayerSigns) uniqueTileSignsKey(column, row int, sign sign) string {
	return fmt.Sprintf("%d:%d:%s", column, row, sign)
}
