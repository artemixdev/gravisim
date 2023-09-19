package physx

const g = 6.6743e-11

func Gravity(mass1, mass2, distance float64) float64 {
	return g * mass1 * mass2 / (distance * distance)
}
