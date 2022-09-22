package noise

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"math/rand"
)

type NoiseMapConfig struct {
	columns, rows          int
	seed                   int64
	frequency, alpha, beta float64
	pointSize, n           int
}

func NewNoiseMapConfig(columns, rows, pointSize, n int, seed int64, frequency, alpha, beta float64) *NoiseMapConfig {
	return &NoiseMapConfig{
		columns:   columns,
		rows:      rows,
		seed:      seed,
		frequency: frequency,
		alpha:     alpha,
		beta:      beta,
		pointSize: pointSize,
		n:         n,
	}
}

type NoiseMap struct {
	Config *NoiseMapConfig
	Points map[int]map[int]float64
}

func (m *NoiseMap) Set(column int, row int, point float64) {
	fetchedColumn, ok := m.Points[column]

	if !ok {
		fetchedColumn = map[int]float64{}
		m.Points[column] = fetchedColumn
	}

	fetchedColumn[row] = point
}

func (m *NoiseMap) Get(column int, row int) float64 {
	if column, ok := m.Points[column]; ok {
		if row, ok := column[row]; ok {
			return row
		}
	}

	panic(fmt.Sprintf("Can't find point by address column [%d] and row [%d]", column, row))
}

func NewNoiseMap(config *NoiseMapConfig) *NoiseMap {
	noise := &NoiseMap{Points: map[int]map[int]float64{}}

	perlinGenerator := perlin.NewPerlinRandSource(
		config.alpha,
		config.beta,
		int32(config.n),
		rand.NewSource(config.seed),
	)

	for column := 0; column < config.columns; column++ {
		for row := 0; row < config.rows; row++ {
			x := float64(column * config.pointSize)
			y := float64(row * config.pointSize)

			point := perlinGenerator.Noise2D(x/config.frequency, y/config.frequency) +
				0.5*perlinGenerator.Noise2D(x/200, y/200) -
				0.2*perlinGenerator.Noise2D(x/100, y/100)

			noise.Set(column, row, point)
		}
	}

	return noise
}
