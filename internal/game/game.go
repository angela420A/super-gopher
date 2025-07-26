package game

import (
	cft "super-gopher/internal/config"
	"super-gopher/internal/player"

	"github.com/hajimehoshi/ebiten/v2"
)

var config = cft.Config

type Game struct {
	Player *player.Player
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.DrawImage(img, nil)
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.Game.ScreenWidth, config.Game.ScreenHeight
}
