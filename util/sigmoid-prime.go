package util

import (
	"github.com/diegodisant/ocrnn-go/matrix"
	"gonum.org/v1/gonum/mat"
)

/**
  Sigmoid: 1 / 1 + e^-x -> sig(x)
  Sigmoid Prime: sig(x) * (1 - sig(x)) -> sigP(x)
*/
func SigmoidPrime(
	matrix mat.Matrix,
	networkOps *matrix.MatrixOperations,
) mat.Matrix {
	rowsLength, _ := matrix.Dims()

	onedArray := make([]float64, rowsLength)

	for matrixIndex := range onedArray {
		onedArray[matrixIndex] = 1
	}

	onedMatrix := mat.NewDense(rowsLength, 1, onedArray)

	// (1 - sig(x))
	sigmoidSubstract := networkOps.Substract(onedMatrix, matrix)

	// sig(x) * (1 - sig(x))
	sigmoidPrime := networkOps.Multiply(onedMatrix, sigmoidSubstract)

	return sigmoidPrime
}
