package storage

type ModelFiles struct {
	HiddenWeights string
	OutputWeights string
}

func NewDefaultModelFiles() ModelFiles {
	return ModelFiles{
		HiddenWeights: HiddenWeightsDefaultFile,
		OutputWeights: OutputWeightsDefaultFile,
	}
}
