package game

import (
	"time"

	"gravisim/app/mathx"
	"gravisim/app/physx"
)

func (self *Game) Update() error {
	now := time.Now()
	defer func() { self.lastUpdate = now }()

	if self.lastUpdate.Unix() == 0 {
		return nil
	}

	delta := now.Sub(self.lastUpdate)
	if delta > time.Second {
		delta = time.Second
	}

	for _, body := range self.bodies {
		body.Position = body.Position.Add(body.Velocity.Mul(delta.Seconds()))
	}

	for _, body := range self.bodies {
		for _, other := range self.bodies {
			if other == body {
				continue
			}

			distanceVector := other.Position.Sub(body.Position)
			distanceScalar := distanceVector.Radius()

			gravityScalar := physx.Gravity(body.Mass, other.Mass, distanceScalar)
			gravityVector := mathx.NewVector2FromPolar(gravityScalar, distanceVector.Angle())

			accelerationVector := gravityVector.Div(body.Mass)
			body.Velocity = body.Velocity.Add(accelerationVector.Mul(delta.Seconds()))
		}
	}

	return nil
}
