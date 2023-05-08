package network

import (
	"gonum.org/v1/gonum/mat"

	"github.com/diegodisant/ocrnn-go/matrix"
	"github.com/diegodisant/ocrnn-go/util"
)

type NeuralNetwork struct {
	inputLayers   int
	hiddenLayers  int
	outputLayers  int
	hiddenWeights *mat.Dense
	outputWeights *mat.Dense
	learningRate  float64
	operations    *matrix.MatrixOperations
}

func NewNeuralNetwork(
	inputLayers,
	hiddenLayers,
	outputLayers int,
	learningRate float64,
) *NeuralNetwork {
	var network = &NeuralNetwork{
		inputLayers:  inputLayers,
		hiddenLayers: hiddenLayers,
		outputLayers: outputLayers,
		learningRate: learningRate,
		operations:   matrix.NewMatrixOperations(),
	}

	network.hiddenWeights = mat.NewDense(
		network.hiddenLayers,
		network.inputLayers,
		util.RandomArray(network.hiddenLayers*network.inputLayers, float64(network.inputLayers)),
	)

	network.outputWeights = mat.NewDense(
		network.outputLayers,
		network.hiddenLayers,
		util.RandomArray(network.outputLayers*network.hiddenLayers, float64(network.hiddenLayers)),
	)

	return network
}
