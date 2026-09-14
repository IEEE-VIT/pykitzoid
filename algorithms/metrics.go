// Package metrics provides common evaluation measures for model predictions.
package metrics

// MeanSquaredError returns the average squared difference between each true
// value and its corresponding prediction.
//
// It panics when the input slices have different lengths because a metric is
// undefined when values cannot be paired. Empty slices return 0.
func MeanSquaredError(trueValues, predictedValues []float64) float64 {
	if len(trueValues) != len(predictedValues) {
		panic("metrics: true and predicted values must have the same length")
	}

	if len(trueValues) == 0 {
		return 0
	}

	var sum float64
	for i, trueValue := range trueValues {
		difference := trueValue - predictedValues[i]
		sum += difference * difference
	}

	return sum / float64(len(trueValues))
}

// MeanAbsoluteError returns the average absolute difference between each true
// value and its corresponding prediction.
//
// It panics when the input slices have different lengths because a metric is
// undefined when values cannot be paired. Empty slices return 0.
func MeanAbsoluteError(trueValues, predictedValues []float64) float64 {
	if len(trueValues) != len(predictedValues) {
		panic("metrics: true and predicted values must have the same length")
	}

	if len(trueValues) == 0 {
		return 0
	}

	var sum float64
	for i, trueValue := range trueValues {
		difference := trueValue - predictedValues[i]
		if difference < 0 {
			difference = -difference
		}
		sum += difference
	}

	return sum / float64(len(trueValues))
}

// MSE is a concise alias for MeanSquaredError.
func MSE(trueValues, predictedValues []float64) float64 {
	return MeanSquaredError(trueValues, predictedValues)
}

// MAE is a concise alias for MeanAbsoluteError.
func MAE(trueValues, predictedValues []float64) float64 {
	return MeanAbsoluteError(trueValues, predictedValues)
}
