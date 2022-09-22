package layers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"zone/core/resources"
	"zone/core/utils/random"
)

func getRandomGroundSprite() *ebiten.Image {
	switch random.Between(0, 1) {
	case 0:
		return resources.GetSpriteOrPanic("ground:grass:1").Original()
	case 1:
		return resources.GetSpriteOrPanic("ground:grass:2").Original()
	}

	return nil
}

func getMountainSpritesheet() *resources.SpriteSheet {
	return resources.GetSpriteSheetOrPanic("mountain:sandstone")
}

func getSingleWaterTileSprite() *ebiten.Image {
	return resources.GetSpriteOrPanic("water:single").Original()
}

func getRandomWaterSprite() *ebiten.Image {
	switch random.Between(0, 1) {
	case 0:
		return resources.GetSpriteOrPanic("water:1").Original()
	case 1:
		return resources.GetSpriteOrPanic("water:2").Original()
	}

	return nil
}

func getCoastlineSpritesheet() *resources.SpriteSheet {
	return resources.GetSpriteSheetOrPanic("coastline")
}

func getRandomSandstoneSmall() *ebiten.Image {
	switch random.Between(0, 1) {
	case 0:
		return resources.GetSpriteOrPanic("ground:grass:sandstone:small:1").Original()
	case 1:
		return resources.GetSpriteOrPanic("ground:grass:sandstone:small:2").Original()
	}

	return nil
}

func getRandomSandstoneHuge() *ebiten.Image {
	switch random.Between(0, 1) {
	case 0:
		return resources.GetSpriteOrPanic("ground:grass:sandstone:huge:1").Original()
	case 1:
		return resources.GetSpriteOrPanic("ground:grass:sandstone:huge:2").Original()
	}

	return nil
}

//func getRandomTree() (*ebiten.Image, int) {
//	switch random.Between(0, 4) {
//	case 0:
//		return resources.GetSpriteOrPanic("tree:dead:small:1").Original(), 0
//	case 1:
//		return resources.GetSpriteOrPanic("tree:oak:small:1").Original(), 0
//	case 2:
//		return resources.GetSpriteOrPanic("tree:oak:small:2").Original(), 0
//	case 3:
//		return resources.GetSpriteOrPanic("tree:pine:small:1").Original(), 0
//	case 4:
//		return resources.GetSpriteOrPanic("tree:pine:small:2").Original(), 0
//	}
//
//	return nil, 0
//}

func getRandomTree() (*ebiten.Image, int) {
	switch random.Between(0, 8) {
	case 0:
		return resources.GetSpriteOrPanic("tree:dead:1").Original(), 16
	case 1:
		return resources.GetSpriteOrPanic("tree:dead:small:1").Original(), 0
	case 2:
		return resources.GetSpriteOrPanic("tree:oak:1").Original(), 16
	case 3:
		return resources.GetSpriteOrPanic("tree:oak:2").Original(), 16
	case 4:
		return resources.GetSpriteOrPanic("tree:oak:small:1").Original(), 0
	case 5:
		return resources.GetSpriteOrPanic("tree:oak:small:2").Original(), 0
	case 6:
		return resources.GetSpriteOrPanic("tree:pine:1").Original(), 16
	case 7:
		return resources.GetSpriteOrPanic("tree:pine:2").Original(), 16
	case 8:
		return resources.GetSpriteOrPanic("tree:pine:small:1").Original(), 0
	case 9:
		return resources.GetSpriteOrPanic("tree:pine:small:2").Original(), 0
	}

	return nil, 0
}

func getRandomFlower() *ebiten.Image {
	switch random.Between(0, 3) {
	case 0:
		return resources.GetSpriteOrPanic("flower:1").Original()
	case 1:
		return resources.GetSpriteOrPanic("flower:2").Original()
	case 2:
		return resources.GetSpriteOrPanic("flower:3").Original()
	case 3:
		return resources.GetSpriteOrPanic("flower:4").Original()
	}

	return nil
}

func getRandomGrass() *ebiten.Image {
	switch random.Between(0, 1) {
	case 0:
		return resources.GetSpriteOrPanic("grass:1").Original()
	case 1:
		return resources.GetSpriteOrPanic("grass:2").Original()
	}

	return nil
}
