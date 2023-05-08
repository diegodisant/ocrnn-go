package util

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func RandomArray(size int, refValue float64) []float64 {
	refValueSqrt := math.Sqrt(refValue)

	uniformDistribution := distuv.Uniform{
		Min: -1 / refValueSqrt,
		Max: 1 / refValueSqrt,
	}

	randomArray := make([]float64, size)

	for arrayIndex := 0; arrayIndex < size; arrayIndex += 1 {
		randomArray[arrayIndex] = uniformDistribution.Rand()
	}

	return randomArray
}
