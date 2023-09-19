package model

import (
	"image/color"

	"gravisim/app/mathx"
)

type Body struct {
	Position mathx.Vector2
	Velocity mathx.Vector2
	Mass     float64
	Radius   float64
	Color    color.Color
}
