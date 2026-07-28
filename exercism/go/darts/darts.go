package darts

import (
	"math"
)

func Score(x, y float64) int {
	var d float64 = math.Sqrt(x*x + y*y)

	if d > 10 {
		return 0
	}

	if d > 5 {
		return 1
	}

	if d > 1 {
		return 5
	}

	return 10
}
