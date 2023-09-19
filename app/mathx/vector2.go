package mathx

import "math"

type Vector2 struct {
	x, y float64
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{x: x, y: y}
}

func NewVector2FromPolar(radius, angle float64) Vector2 {
	return Vector2{x: radius * math.Cos(angle), y: radius * math.Sin(angle)}
}

func (self Vector2) X() float64 {
	return self.x
}

func (self Vector2) Y() float64 {
	return self.y
}

func (self Vector2) Radius() float64 {
	return math.Sqrt(self.x*self.x + self.y*self.y)
}

func (self Vector2) Angle() float64 {
	if self.x > 0 {
		angle := math.Atan(self.y / self.x)
		if angle < 0 {
			angle += 2 * math.Pi
		}
		return angle
	}

	if self.x < 0 {
		angle := math.Atan(self.y/self.x) + math.Pi
		if angle < 0 {
			angle += 2 * math.Pi
		}
		return angle
	}

	if self.x == 0 && self.y > 0 {
		return math.Pi / 2
	}

	if self.x == 0 && self.y < 0 {
		return 3 * math.Pi / 2
	}

	return 0 // self.x == 0 && self.y == 0
}

func (self Vector2) Add(other Vector2) Vector2 {
	return Vector2{x: self.x + other.x, y: self.y + other.y}
}

func (self Vector2) Sub(other Vector2) Vector2 {
	return Vector2{x: self.x - other.x, y: self.y - other.y}
}

func (self Vector2) Mul(value float64) Vector2 {
	return Vector2{x: self.x * value, y: self.y * value}
}

func (self Vector2) Div(value float64) Vector2 {
	return Vector2{x: self.x / value, y: self.y / value}
}
