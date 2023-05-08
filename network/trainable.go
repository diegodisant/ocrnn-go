package network

type Trainable interface {
	Train(inputData, targetData []float64)
}
