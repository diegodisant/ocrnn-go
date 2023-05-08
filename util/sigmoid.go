package util

import (
	"math"

	"github.com/diegodisant/ocrnn-go/matrix"
)

var Sigmoid matrix.MatrixApplicable = func(rowIndex, colIndex int, currentValue float64) float64 {
	return 1.0 / (1 + math.Exp(-1*currentValue))
}
