package game

import "github.com/hajimehoshi/ebiten/v2"

func (self *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = int(float64(outsideWidth) * ebiten.DeviceScaleFactor())
	screenHeight = int(float64(outsideHeight) * ebiten.DeviceScaleFactor())
	return
}
