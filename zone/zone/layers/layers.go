package layers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zone/zone/noise"
)

const (
	ReliefLayer = "relief"
	TreeLayer   = "plants"
)

type Layer interface {
	Update() error
	Draw(screen *ebiten.Image)
	SetNoise(noise *noise.NoiseMap)
	GetSigns() *LayerSigns
}

type LayerConfig struct {
	width, height int
	columns, rows int
	tileSize      int
}

func NewLayerConfig(width, height, columns, rows, tileSize int) *LayerConfig {
	return &LayerConfig{width: width, height: height, columns: columns, rows: rows, tileSize: tileSize}
}
