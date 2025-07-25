package assets

import (
	"embed"
	"image"
	_ "image/png"
	logger "super-gopher/utils/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed characters/* backgrounds/*
var assets embed.FS

var PlayerSprite = mustLoadImage("characters/tile_0000.png")

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
