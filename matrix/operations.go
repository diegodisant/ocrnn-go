package matrix

import "gonum.org/v1/gonum/mat"

type MatrixOperations struct{}

func NewMatrixOperations() *MatrixOperations {
	return &MatrixOperations{}
}

var _ MatrixOperable = (*MatrixOperations)(nil)

func (operations *MatrixOperations) Add(firstMatrix, secondMatrix mat.Matrix) mat.Matrix {
	rowsLength, colsLength := firstMatrix.Dims()

	resultMatrix := mat.NewDense(rowsLength, colsLength, nil)

	resultMatrix.Add(firstMatrix, secondMatrix)

	return resultMatrix
}

func (operations *MatrixOperations) Substract(firstMatrix, secondMatrix mat.Matrix) mat.Matrix {
	rowsLength, colsLength := firstMatrix.Dims()

	resultMatrix := mat.NewDense(rowsLength, colsLength, nil)

	resultMatrix.Sub(firstMatrix, secondMatrix)

	return resultMatrix
}

func (operations *MatrixOperations) Dot(firstMatrix, secondMatrix mat.Matrix) mat.Matrix {
	firstMatrixRowsLength, _ := firstMatrix.Dims()
	_, secondMatrixColsLength := secondMatrix.Dims()

	resultMatrix := mat.NewDense(firstMatrixRowsLength, secondMatrixColsLength, nil)

	resultMatrix.Product(firstMatrix, secondMatrix)

	return resultMatrix
}

func (operations *MatrixOperations) Scale(scalar float64, matrix mat.Matrix) mat.Matrix {
	rowsLength, colsLength := matrix.Dims()

	resultMatrix := mat.NewDense(rowsLength, colsLength, nil)

	resultMatrix.Scale(scalar, matrix)

	return resultMatrix
}

func (operations *MatrixOperations) AddScalar(scalar float64, matrix mat.Matrix) mat.Matrix {
	rowsLength, colsLength := matrix.Dims()

	scalarMatrixDim := rowsLength * colsLength

	scalarMatrix := make([]float64, scalarMatrixDim)

	for matrixIndex := 0; matrixIndex < scalarMatrixDim; matrixIndex += 1 {
		scalarMatrix[matrixIndex] = scalar
	}

	resultMatrix := mat.NewDense(rowsLength, colsLength, scalarMatrix)

	return operations.Add(matrix, resultMatrix)
}

func (operations *MatrixOperations) Multiply(firstMatrix, secondMatrix mat.Matrix) mat.Matrix {
	rowsLength, colsLength := firstMatrix.Dims()

	resultMatrix := mat.NewDense(rowsLength, colsLength, nil)

	resultMatrix.MulElem(firstMatrix, secondMatrix)

	return resultMatrix
}

func (operations *MatrixOperations) Apply(
	onTraverseAction MatrixApplicable,
	matrix mat.Matrix,
) mat.Matrix {

}
