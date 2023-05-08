package network

import (
	"github.com/diegodisant/ocrnn-go/util"
	"gonum.org/v1/gonum/mat"
)

var _ Predictable = (*NeuralNetwork)(nil)

func (net *NeuralNetwork) Predict(inputData []float64) mat.Matrix {
	inputs := mat.NewDense(len(inputData), 1, inputData)

	hiddenInputs := net.operations.Dot(net.hiddenWeights, inputs)
	hiddenOutputs := net.operations.Apply(util.Sigmoid, hiddenInputs)
	resultInputs := net.operations.Dot(net.outputWeights, hiddenOutputs)
	resultOutputs := net.operations.Apply(util.Sigmoid, resultInputs)

	return resultOutputs
}
