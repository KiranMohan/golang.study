package kiran

import "math"

const maxDelta float64 = 0.0001

func Sqrt(x float64) float64 {
	var z = x / 2.0
	//fmt.Printf("initial value %g\n", z)
	for delta := 1.0; delta > maxDelta; {
		oldZ := z
		z = z - (z*z-x)/(2*x)
		delta = math.Abs(z - oldZ)
	}
	return z
}
