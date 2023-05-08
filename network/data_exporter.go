package network

import (
	"os"

	"github.com/diegodisant/ocrnn-go/storage"
)

var _ storage.Exportable = (*NeuralNetwork)(nil)

func (net *NeuralNetwork) Export(model storage.ModelFiles) {
	hiddenWeightsFile, err := os.Create(model.HiddenWeights)

	if err != nil {
		// @TODO: handle logging correctly

		return
	}

	defer hiddenWeightsFile.Close()

	net.hiddenWeights.MarshalBinaryTo(hiddenWeightsFile)

	outputWeightsFile, err := os.Create(model.OutputWeights)

	if err != nil {
		// @TODO: handle logging correctly

		return
	}

	defer outputWeightsFile.Close()

	net.outputWeights.MarshalBinaryTo(outputWeightsFile)
}
