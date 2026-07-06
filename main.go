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
)

type player struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func (p player) MinX() float64 { return p.x }
func (p player) MaxX() float64 { return p.x + p.width }
func (p player) MinY() float64 { return p.y }
func (p player) MaxY() float64 { return p.y + p.height }

func (p *player) Control() {

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.x -= 3.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.x += 3.0
	}
	if p.x < 0 {
		p.x = 0
	}
	maxLinesX := screenWidth - p.width
	if maxLinesX < p.x {
		p.x = maxLinesX
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

type Game struct {
	p player
	b ball
}

func (g *Game) Update() error {
	g.p.Control()
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
			vy: 0,
		},
	}
	ebiten.SetWindowSize(720, 960)
	ebiten.SetWindowTitle("ブロック崩し")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
