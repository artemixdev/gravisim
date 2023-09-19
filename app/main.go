package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"gravisim/app/game"
	"gravisim/app/mathx"
	"gravisim/app/model"
)

func main() {
	g := game.New()

	g.Spawn(model.Body{
		Position: mathx.NewVector2(-2, 0),
		Velocity: mathx.NewVector2(0, -0.6),
		Mass:     1e11,
		Radius:   0.5,
		Color:    colornames.Lightgreen,
	})

	g.Spawn(model.Body{
		Position: mathx.NewVector2(2, 0),
		Velocity: mathx.NewVector2(0, 0.6),
		Mass:     1e11,
		Radius:   0.5,
		Color:    colornames.Skyblue,
	})

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("gravisim")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
