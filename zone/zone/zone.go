package zone

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
	"zone/core/tiles"
	"zone/zone/noise"
	"zone/zone/zone/layers"
)

type Zone struct {
	tiles            *tiles.Collection
	layers           []layers.Layer
	layersIndexByKey map[string]int
}

func NewZone(width, height, columns, rows, tileSize int) *Zone {
	region := &Zone{layers: []layers.Layer{}, layersIndexByKey: map[string]int{}}
	region.tiles = tiles.NewEmptyTileCollection(columns, rows)

	noiseConfig := noise.NewNoiseMapConfig(columns, rows, 16, 3, 10, 200, 3, 3)
	noiseMap := noise.NewNoiseMap(noiseConfig)

	layerConfig := layers.NewLayerConfig(width, height, columns, rows, tileSize)
	_ = region.AddLayer(layers.ReliefLayer, layers.NewRelief(layerConfig, noiseMap, region.tiles))
	_ = region.AddLayer(layers.TreeLayer, layers.NewPlants(layerConfig, noiseMap, region.tiles))

	return region
}

func (z *Zone) AddLayer(alias string, layer layers.Layer) error {
	if _, ok := z.layersIndexByKey[alias]; ok {
		return fmt.Errorf("layer with alias [%s] already added", alias)
	}

	z.layers = append(z.layers, layer)
	index := len(z.layers) - 1
	z.layersIndexByKey[alias] = index

	return nil
}

func (z *Zone) GetLayer(alias string) (layers.Layer, error) {
	index, ok := z.layersIndexByKey[alias]

	if !ok {
		return nil, fmt.Errorf("layer with alias [%s] not found", alias)
	}

	return z.layers[index], nil
}

func (z *Zone) Update() error {
	for _, layer := range z.layers {
		err := layer.Update()
		if err != nil {
			log.Print(err)
		}
	}

	return nil
}

func (z *Zone) Draw(screen *ebiten.Image) {
	for _, layer := range z.layers {
		layer.Draw(screen)
	}
}
