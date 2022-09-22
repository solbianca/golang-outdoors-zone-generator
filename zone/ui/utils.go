package ui

import (
	"image/color"
	"math/rand"
	"time"
	"zone/core/utils"
)

var (
	randomGenerator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

func hexToColor(h string) color.Color {
	return utils.HexToColor(h)
}
