package network

import (
	"github.com/diegodisant/ocrnn-go/util"
	"gonum.org/v1/gonum/mat"
)

var _ Predictable = (*NeuralNetwork)(nil)

func (network *NeuralNetwork) Predict(inputData []float64) mat.Matrix {
	inputs := mat.NewDense(len(inputData), 1, inputData)
  
	hiddenInputs := network.operations.Dot(network.hiddenWeights, inputs)
	hiddenOutputs := network.operations.Apply(util.Sigmoid, hiddenInputs)
	resultInputs := network.operations.Dot(network.outputWeights, hiddenOutputs)
	resultOutputs := network.operations.Apply(util.Sigmoid, resultInputs)

	return resultOutputs
}
