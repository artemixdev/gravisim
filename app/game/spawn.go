package game

import "gravisim/app/model"

func (self *Game) Spawn(bodies ...model.Body) {
	for _, body := range bodies {
		self.bodies = append(self.bodies, &body)
	}
}
