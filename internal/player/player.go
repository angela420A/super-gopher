package player

import (
	ast "super-gopher/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

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
