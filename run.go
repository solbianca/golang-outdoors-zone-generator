package main

import (
	"log"
	"zone/core/resources"
	"zone/zone"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(zone.ScreenWidth, zone.ScreenHeight)
	ebiten.SetWindowTitle("Arena")
	ebiten.SetScreenClearedEveryFrame(true)

	loadResources()

	game := zone.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func loadResources() {
	resources.LoadFont("mana", "resources/fonts/manaspace.ttf")
	resources.LoadFont("pixel", "resources/fonts/pixel-nes.otf")
	resources.LoadFont("roboto", "resources/fonts/RobotoMono-SemiBold.ttf")
	resources.LoadFont("noto", "resources/fonts/NotoSans-Regular.ttf")
	resources.LoadFont("notobold", "resources/fonts/NotoSans-Bold.ttf")

	resources.LoadSprite("liteui:input", "resources/liteui/Input Field 1.png")
	resources.LoadSprite("liteui:btn:dark", "resources/liteui/Big Dark Btn.png")
	resources.LoadSprite("liteui:btn:circular", "resources/liteui/Circular Btn.png")
	resources.LoadSprite("liteui:btn:radio", "resources/liteui/Radio Btn Frame.png")
	resources.LoadSprite("liteui:bar-frame", "resources/liteui/Bar Frame.png")

	resources.LoadSprite("liteui:panel", "resources/liteui/panel.png")

	resources.LoadSprite("ui:button:idle", "resources/ui/button-idle.png")
	resources.LoadSprite("ui:button:hover", "resources/ui/button-hover.png")
	resources.LoadSprite("ui:button:pressed", "resources/ui/button-pressed.png")
	resources.LoadSprite("ui:button:active", "resources/ui/button-active.png")

	resources.LoadSprite("ground:grass:1", "resources/sprites/ground_grass_1.png")
	resources.LoadSprite("ground:grass:2", "resources/sprites/ground_grass_2.png")
	resources.LoadSprite("ground:grass:sandstone:small:1", "resources/sprites/ground_grass_sandstone_small_1.png")
	resources.LoadSprite("ground:grass:sandstone:small:2", "resources/sprites/ground_grass_sandstone_small_2.png")
	resources.LoadSprite("ground:grass:sandstone:huge:1", "resources/sprites/ground_grass_sandstone_huge_1.png")
	resources.LoadSprite("ground:grass:sandstone:huge:2", "resources/sprites/ground_grass_sandstone_huge_2.png")
	resources.LoadSpriteSheet("mountain:sandstone", "resources/sprites/mountain_sandstone_spritesheet.png", 16, 16)
	resources.LoadSprite("water:1", "resources/sprites/water_1.png")
	resources.LoadSprite("water:2", "resources/sprites/water_2.png")
	resources.LoadSprite("water:single", "resources/sprites/water_single.png")
	resources.LoadSpriteSheet("coastline", "resources/sprites/coastline_spritesheet.png", 16, 16)

	resources.LoadSprite("flower:1", "resources/sprites/flower_1.png")
	resources.LoadSprite("flower:2", "resources/sprites/flower_2.png")
	resources.LoadSprite("flower:3", "resources/sprites/flower_3.png")
	resources.LoadSprite("flower:4", "resources/sprites/flower_4.png")
	resources.LoadSprite("grass:1", "resources/sprites/grass_1.png")
	resources.LoadSprite("grass:2", "resources/sprites/grass_2.png")

	resources.LoadSprite("tree:dead:1", "resources/sprites/tree_dead_1.png")
	resources.LoadSprite("tree:dead:small:1", "resources/sprites/tree_dead_small_1.png")
	resources.LoadSprite("tree:oak:1", "resources/sprites/tree_oak_1.png")
	resources.LoadSprite("tree:oak:2", "resources/sprites/tree_oak_2.png")
	resources.LoadSprite("tree:oak:small:1", "resources/sprites/tree_oak_small_1.png")
	resources.LoadSprite("tree:oak:small:2", "resources/sprites/tree_oak_small_2.png")
	resources.LoadSprite("tree:pine:1", "resources/sprites/tree_pine_1.png")
	resources.LoadSprite("tree:pine:2", "resources/sprites/tree_pine_2.png")
	resources.LoadSprite("tree:pine:small:1", "resources/sprites/tree_pine_small_1.png")
	resources.LoadSprite("tree:pine:small:2", "resources/sprites/tree_pine_small_2.png")
}
