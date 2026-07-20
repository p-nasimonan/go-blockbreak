package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 240.0
	screenHeight = 320.0
	maxPlayerV   = 3.0
	maxBallV     = 3.0
)

type player struct {
	x      float64
	y      float64
	vx     float64
	vy     float64
	width  float64
	height float64
}

func (p player) MinX() float64 { return p.x }
func (p player) MaxX() float64 { return p.x + p.width }
func (p player) MinY() float64 { return p.y }
func (p player) MaxY() float64 { return p.y + p.height }
func (b ball) MinX() float64   { return b.x - b.r }
func (b ball) MaxX() float64   { return b.x + b.r }
func (b ball) MinY() float64   { return b.y - b.r }
func (b ball) MaxY() float64   { return b.y + b.r }

func (p *player) Control() {

	if ebiten.IsKeyPressed(ebiten.KeyH) {
		p.vx += -0.6
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		p.vx += 0.6
	}

	if ebiten.IsKeyPressed(ebiten.KeyK) {
		p.vy += -0.6
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		p.vy += 0.6
	}

	if p.x < 0 {
		p.x = 0
	}
	maxLinesX := screenWidth - p.width
	if maxLinesX < p.x {
		p.x = maxLinesX
	}

	if p.y < 0 {
		p.y = 0
	}
	maxLinesY := screenHeight - p.height
	if maxLinesY < p.y {
		p.y = maxLinesY
	}

}

func (p *player) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(p.x),
		float32(p.y),
		float32(p.width),
		float32(p.height),
		color.RGBA{100, 180, 255, 255},
		false,
	)
}

func (p *player) Update() {
	p.x += p.vx
	p.y += p.vy

	p.vx *= 0.7
	p.vy *= 0.7
	if p.vx > maxPlayerV {
		p.vx = maxPlayerV
	}
	if p.vx < -maxPlayerV {
		p.vx = -maxPlayerV
	}
	if p.vy > maxPlayerV {
		p.vy = maxPlayerV
	}
	if p.vy < -maxPlayerV {
		p.vy = -maxPlayerV
	}
}

type ball struct {
	x  float64
	y  float64
	r  float64
	vx float64
	vy float64
}

func (b *ball) Draw(screen *ebiten.Image) {
	vector.FillCircle(
		screen,
		float32(b.x),
		float32(b.y),
		float32(b.r),
		color.RGBA{255, 120, 120, 255},
		false,
	)
}

func (b *ball) Update(p *player) {
	b.x += b.vx
	b.y += b.vy
	if b.MinX() < 0 {
		b.vx = -b.vx
		b.x = b.r
	}
	if b.MaxX() > screenWidth {
		b.vx = -b.vx
		b.x = screenWidth - b.r
	}
	if b.MinY() < 0 {
		b.vy = -b.vy
		b.y = b.r
	}
	if b.MaxY() > screenHeight {
		b.vy = -b.vy
		b.y = screenHeight - b.r
	}

	isColliding := b.MinX() < p.MaxX() && b.MaxX() > p.MinX() && b.MinY() < p.MaxY() && b.MaxY() > p.MinY()
	if isColliding {
		if b.vy > 0 {
			b.y = p.y - b.r
		} else {
			b.y = p.y + p.height + b.r
		}

		b.vy = -b.vy
		playerCenterX := p.x + p.width/2
		b.vx = 2*(b.x-playerCenterX)/p.width + (p.vx / 2)

	}
}

type Game struct {
	p player
	b ball
}

func (g *Game) Update() error {
	g.p.Control()
	g.p.Update()
	g.b.Update(&g.p)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})
	g.p.Draw(screen)
	g.b.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
	w = int(screenWidth)
	h = int(screenHeight)
	return w, h
}

func main() {
	initialWidth := 50.0
	initialHeight := 12.0

	game := &Game{
		p: player{
			x:      (screenWidth - initialWidth) / 2,
			y:      screenHeight - 30,
			width:  initialWidth,
			height: initialHeight,
		},
		b: ball{
			x:  (screenWidth - initialWidth) / 2,
			y:  30,
			r:  5,
			vx: 0,
			vy: 2,
		},
	}
	ebiten.SetWindowSize(720, 960)
	ebiten.SetWindowTitle("ブロック崩し")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
