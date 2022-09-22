package assets

import (
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	images *imageStorage
)

func init() {
	images = newImageStorage()
}

type imageStorage struct {
	images map[string]*ebiten.Image
}

func newImageStorage() *imageStorage {
	return &imageStorage{images: map[string]*ebiten.Image{}}
}

func GetImage(alias string) (*ebiten.Image, error) {
	if img, ok := images.images[alias]; ok {
		return img, nil
	}

	return nil, fmt.Errorf("image by alias [%s] not loaded", alias)
}

func LoadImage(alias, path string) {
	file, err := ebitenutil.OpenFile(path)

	if err != nil {
		panic(fmt.Errorf("can't open file by path [%s], %v", path, err))
	}
	defer func() {
		_ = file.Close()
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(fmt.Errorf("can't decode file by path [%s], %v", path, err))
	}

	images.images[alias] = ebiten.NewImageFromImage(img)
}
