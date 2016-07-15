package kiran

import (
	"fmt"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	for i := 4; i < 256; i++ {
		fmt.Printf("Sqrt of %d - Expected value : %f, Actual Value : %f\n", i, math.Sqrt(float64(i)), Sqrt(float64(i)))
	}
	t.Errorf("Kiran says it failed")
}
