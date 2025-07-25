package main

import (
	"log"
	"math"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	ast "super-gopher/assets"
	_ "super-gopher/utils/config"
	_ "super-gopher/utils/logger"
)

type Vector struct {
	X float64
	Y float64
}

// Player
type Player struct {
	position Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer() *Player {
	sprite := ast.PlayerSprite

	return &Player{
		position: Vector{X: 100, Y: 100},
		sprite:   sprite,
		rotation: 0,
	}
}

func (p *Player) Update() {
	// speed := math.Pi / float64(ebiten.TPS())
	speed := 5.0
	turn := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.rotation -= turn
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.rotation += turn
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

// Game
const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/backgrounds/tile_0000.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	g := &Game{
		player: NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
