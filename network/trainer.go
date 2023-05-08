package network

import (
	"github.com/diegodisant/ocrnn-go/util"
	"gonum.org/v1/gonum/mat"
)

var _ Trainable = (*NeuralNetwork)(nil)

func (net *NeuralNetwork) Train(inputData, targetData []float64) {
	inputs := mat.NewDense(len(inputData), 1, inputData)

	hiddenInputs := net.operations.Dot(net.hiddenWeights, inputs)
	hiddenOutputs := net.operations.Apply(util.Sigmoid, hiddenInputs)
	resultInputs := net.operations.Dot(net.outputWeights, hiddenOutputs)
	finalOutputs := net.operations.Apply(util.Sigmoid, resultInputs)

	targets := mat.NewDense(len(targetData), 1, targetData)
	outputErrors := net.operations.Substract(targets, finalOutputs)
	hiddenErrors := net.operations.Dot(net.outputWeights.T(), outputErrors)

	/**
	  oW: output weights
	  lR: learning rate
	  Ek: output errors
	  sigP: sigmoid prime
	  hO: hidden output
	  (*): dot operation

	  oW = oW + (lR * ((Ek * sigP) (*) hO))
	*/

	outputSigPErrorsMult := net.operations.Multiply(
		outputErrors,
		util.SigmoidPrime(finalOutputs, net.operations),
	)

	hiddenOutputsWithSigPErrors := net.operations.Dot(
		outputSigPErrorsMult,
		hiddenOutputs.T(),
	)

	outputLearningRateScalation := net.operations.Scale(
		net.learningRate,
		hiddenOutputsWithSigPErrors,
	)

	net.outputWeights = net.operations.Add(
		net.outputWeights,
		outputLearningRateScalation,
	).(*mat.Dense)

	/**
	  oW: hidden weights
	  lR: learning rate
	  Ek: hidden errors
	  sigP: sigmoid prime
	  i: inputs
	  (*): dot operation

	  oW = oW + (lR * ((Ek * sigP) (*) i))
	*/

	hiddenSigPErrorsMult := net.operations.Multiply(
		hiddenErrors,
		util.SigmoidPrime(hiddenOutputs, net.operations),
	)

	inputWithSigPErrors := net.operations.Dot(
		hiddenSigPErrorsMult,
		inputs.T(),
	)

	hiddenLearningRateScalation := net.operations.Scale(
		net.learningRate,
		inputWithSigPErrors,
	)

	net.hiddenWeights = net.operations.Add(
		net.hiddenWeights,
		hiddenLearningRateScalation,
	).(*mat.Dense)
}
