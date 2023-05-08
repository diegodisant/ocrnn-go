package network

import "gonum.org/v1/gonum/mat"

type Predictable interface {
	Predict(inputData []float64) mat.Matrix
}
