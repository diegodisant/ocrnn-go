package network

import (
	"os"

	"github.com/diegodisant/ocrnn-go/storage"
)

var _ storage.Importable = (*NeuralNetwork)(nil)

func (net *NeuralNetwork) Import(model storage.ModelFiles) {
	hiddenWeightsFile, err := os.Create(model.HiddenWeights)

	if err != nil {
		// @TODO: handle logging correctly

		return
	}

	defer hiddenWeightsFile.Close()

	net.hiddenWeights.Reset()
	net.hiddenWeights.UnmarshalBinaryFrom(hiddenWeightsFile)

	outputWeightsFile, err := os.Create(model.OutputWeights)

	if err != nil {
		// @TODO: handle logging correctly

		return
	}

	defer outputWeightsFile.Close()

	net.outputWeights.Reset()
	net.outputWeights.UnmarshalBinaryFrom(outputWeightsFile)
}
