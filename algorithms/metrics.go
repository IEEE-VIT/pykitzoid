package algorithms

// MeanSquaredError calculates the average squared difference
// between the true and predicted values.
func MeanSquaredError(trueValues, predictedValues []float64) float64 {
	if len(trueValues) != len(predictedValues) || len(trueValues) == 0 {
		return 0
	}

	var sum float64

	for i := range trueValues {
		diff := trueValues[i] - predictedValues[i]
		sum += diff * diff
	}

	return sum / float64(len(trueValues))
}

// MeanAbsoluteError calculates the average absolute difference
// between the true and predicted values.
func MeanAbsoluteError(trueValues, predictedValues []float64) float64 {
	if len(trueValues) != len(predictedValues) || len(trueValues) == 0 {
		return 0
	}

	var sum float64

	for i := range trueValues {
		diff := trueValues[i] - predictedValues[i]

		if diff < 0 {
			diff = -diff
		}

		sum += diff
	}

	return sum / float64(len(trueValues))
}