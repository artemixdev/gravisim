package mathx

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVector2_Angle(t *testing.T) {
	v := Vector2{x: 0, y: 0} // center
	require.Equal(t, 0.000, trunc3(v.Angle()))

	v = Vector2{x: 1, y: 0} // 0
	require.Equal(t, trunc3(0*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: 1, y: 1} // 45
	require.Equal(t, trunc3(1*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: 0, y: 1} // 90
	require.Equal(t, trunc3(2*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: -1, y: 1} // 135
	require.Equal(t, trunc3(3*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: -1, y: 0} // 180
	require.Equal(t, trunc3(4*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: -1, y: -1} // 225
	require.Equal(t, trunc3(5*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: 0, y: -1} // 270
	require.Equal(t, trunc3(6*math.Pi/4), trunc3(v.Angle()))

	v = Vector2{x: 1, y: -1} // 315
	require.Equal(t, trunc3(7*math.Pi/4), trunc3(v.Angle()))
}

func trunc3(value float64) float64 {
	return math.Trunc(value*1000) / 1000
}
