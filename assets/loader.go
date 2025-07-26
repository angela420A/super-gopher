package assets

import (
	"embed"
	"image"
	_ "image/png"
	logger "super-gopher/internal/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images
var assets embed.FS

var PlayerSprite = mustLoadImage("images/characters/tile_0000.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		logger.Error("Failed to open image file", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		logger.Error("Failed to decode image", err)
	}
	return ebiten.NewImageFromImage(img)
}
