package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	_ "super-gopher/internal/config"
	game "super-gopher/internal/game"
	_ "super-gopher/internal/logger"
	ply "super-gopher/internal/player"
)

type Vector struct {
	X float64
	Y float64
}

// Game
func main() {
	g := &game.Game{
		Player: ply.NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
