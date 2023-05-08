package util

import "math"

func Sigmoid(rowIndex, colIndex int, currentValue float64) float64 {
	return 1.0 / (1 + math.Exp(-1*currentValue))
}
