package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"gravisim/app/mathx"
)

func (self *Game) Draw(screen *ebiten.Image) {
	for _, body := range self.bodies {
		x, y := convertVector(body.Position, screen.Bounds().Dx(), screen.Bounds().Dy())
		r := convertDistance(body.Radius)
		vector.DrawFilledCircle(screen, x, y, r, body.Color, true)
	}

	image := ebiten.NewImage(100, 100)
	ebitenutil.DebugPrint(image, fmt.Sprintf("%.0f", ebiten.ActualFPS()))

	factor := ebiten.DeviceScaleFactor()
	geom := ebiten.GeoM{}
	geom.Scale(1.5*factor, 1.5*factor)

	screen.DrawImage(image, &ebiten.DrawImageOptions{GeoM: geom})
}

// pixels per unit - how many pixels are in range from 0 to 1 for a vector
var ppu = 35 * ebiten.DeviceScaleFactor()

func convertVector(value mathx.Vector2, screenWidth, screenHeight int) (x, y float32) {
	x = float32(screenWidth)/2 + convertDistance(value.X())
	y = float32(screenHeight)/2 - convertDistance(value.Y())
	return
}

func convertDistance(distance float64) float32 {
	return float32(distance * ppu)
}
