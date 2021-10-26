package atom

import (
	"math"
)

func Min(a,b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a,b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func Sign(a float64) float64 {
	if a >= 0 {
		return 1
	} else {
		return -1
	}
}

func Precision(a float64) int {
	return int(math.Log10(1/a))
}

func Round(a float64, p int) float64 {
	time := math.Pow10(p)
	return float64(int64(a*time))/time
}