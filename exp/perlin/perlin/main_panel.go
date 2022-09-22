package perlin

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	alpha                = 2.0
	beta                 = 2.0
	n                    = 3
	frequency            = 400.0
	seed           int64 = 999
	pointSize            = 16
	isStateChanged       = true
)

type MainPanel struct {
	Panel      *Panel
	noise      *noiseMap
	noiseImage *ebiten.Image
}

func NewMainPanel(windowWidth, windowHeight int, globalX, globalY float64, width, height int) *MainPanel {
	panel := NewPanel(windowWidth, windowHeight, globalX, globalY, width, height)

	mainPanel := &MainPanel{Panel: panel}
	mainPanel.generateNoiseMapImage()

	return mainPanel
}

func (p *MainPanel) GetSeed() int64 {
	return seed
}

func (p *MainPanel) SetSeed(newSeed int64) {
	isStateChanged = true
	seed = newSeed
}

func (p *MainPanel) GetPointSize() int {
	return pointSize
}

func (p *MainPanel) SetPointSize(newPointSize int) {
	isStateChanged = true
	pointSize = newPointSize
}

func (p *MainPanel) GetFrequency() float64 {
	return frequency
}

func (p *MainPanel) GetAlpha() float64 {
	return alpha
}

func (p *MainPanel) SetAlpha(newAlpha float64) {
	isStateChanged = true
	alpha = newAlpha
}

func (p *MainPanel) GetBeta() float64 {
	return beta
}

func (p *MainPanel) SetBeta(newBeta float64) {
	isStateChanged = true
	beta = newBeta
}

func (p *MainPanel) GetN() int {
	return n
}

func (p *MainPanel) SetN(newN int) {
	isStateChanged = true
	n = newN
}

func (p *MainPanel) SetFrequency(newFrequency float64) {
	isStateChanged = true
	frequency = newFrequency
}

func (p *MainPanel) IsStateChanged() bool {
	return isStateChanged
}

func (p *MainPanel) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0), float64(0))

	screen.DrawImage(p.noiseImage, op)
}

func (p *MainPanel) Update() {
	if isStateChanged == false {
		return
	}

	p.generateNoiseMapImage()

	isStateChanged = false
}

func (p *MainPanel) generateNoiseMapImage() {
	p.noise = createNoiseMap(p)
	p.noiseImage = drawNoiseImage(p)
}

func createNoiseMap(p *MainPanel) *noiseMap {
	noise := &noiseMap{map[int]map[int]float64{}}

	perlinGenerator := perlin.NewPerlinRandSource(alpha, beta, int32(n), rand.NewSource(seed))
	columns := p.Panel.Width / pointSize
	rows := p.Panel.Width / pointSize

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			x := float64(column * pointSize)
			y := float64(row * pointSize)

			point := perlinGenerator.Noise2D(x/frequency, y/frequency) +
				0.5*perlinGenerator.Noise2D(x/200, y/200) -
				0.2*perlinGenerator.Noise2D(x/100, y/100)

			noise.set(column, row, point)
		}
	}

	return noise
}

func drawNoiseImage(p *MainPanel) *ebiten.Image {
	noiseImage := ebiten.NewImage(p.Panel.Width, p.Panel.Height)

	columns := p.Panel.Width / pointSize
	rows := p.Panel.Height / pointSize

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			x := float64(column * pointSize)
			y := float64(row * pointSize)

			x, y = p.Panel.LocalToGlobalCoordinates(x, y)

			point := p.noise.get(column, row)

			rectColor := color.RGBA{255, 0, 0, 0}
			if point >= -1 && point < -0.66 {
				rectColor = color.RGBA{0, 0, 0, 255}
			} else if point >= -0.66 && point < -0.33 {
				rectColor = color.RGBA{60, 60, 60, 255}
			} else if point >= -0.33 && point < 0 {
				rectColor = color.RGBA{120, 120, 120, 255}
			} else if point >= 0 && point < 0.33 {
				rectColor = color.RGBA{180, 180, 180, 255}
			} else if point >= 0.33 && point < 0.66 {
				rectColor = color.RGBA{240, 240, 240, 255}
			} else if point >= 0.66 && point <= 1 {
				rectColor = color.RGBA{255, 255, 255, 255}
			}

			ebitenutil.DrawRect(noiseImage, x, y, float64(pointSize), float64(pointSize), rectColor)
		}
	}

	return noiseImage
}

type noiseMap struct {
	points map[int]map[int]float64
}

func (m *noiseMap) set(column int, row int, point float64) {
	fetchedColumn, ok := m.points[column]

	if !ok {
		fetchedColumn = map[int]float64{}
		m.points[column] = fetchedColumn
	}

	fetchedColumn[row] = point
}

func (m *noiseMap) get(column int, row int) float64 {
	if column, ok := m.points[column]; ok {
		if row, ok := column[row]; ok {
			return row
		}
	}

	panic(fmt.Sprintf("Can't find point by address column [%d] and row [%d]", column, row))
}
