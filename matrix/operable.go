package matrix

import "gonum.org/v1/gonum/mat"

type MatrixOperable interface {
	Add(firstMatrix, secondMatrix mat.Matrix) mat.Matrix

	Substract(firstMatrix, secondMatrix mat.Matrix) mat.Matrix

	Dot(firstMatrix, secondMatrix mat.Matrix) mat.Matrix

	Scale(scalar float64, matrix mat.Matrix) mat.Matrix

	AddScalar(scalar float64, matrix mat.Matrix) mat.Matrix

	Multiply(firstMatrix, secondMatrix mat.Matrix) mat.Matrix

	Apply(onTraverseAction MatrixApplicable, matrix mat.Matrix) mat.Matrix
}
