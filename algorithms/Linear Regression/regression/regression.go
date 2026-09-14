// Package regression contains the numerical operations shared by the demo and API.
package regression

func MeanOfColumn(arr [][]float32, col_number int) float32 {
	// calculate the sum of the array
	var calculated_mean_value float32
	calculated_mean_value = 0
	for i := 0; i < len(arr); i++ {
		calculated_mean_value += arr[i][col_number]
	}
	calculated_mean_value = calculated_mean_value / float32(len(arr))
	return calculated_mean_value
}

func CalculateSlopeAndIntercept(csv_object [][]float32) (float32, float32) {
	var numerator, denominator float32
	numerator = 0
	denominator = 0

	// calculate the mean of x and y
	var x_mean, y_mean float32
	x_mean = MeanOfColumn(csv_object, 0)
	y_mean = MeanOfColumn(csv_object, 1)

	// calculate the numerator and denominator
	for i := 0; i < len(csv_object); i++ {
		numerator += (csv_object[i][0] - x_mean) * (csv_object[i][1] - y_mean)
		denominator += (csv_object[i][0] - x_mean) * (csv_object[i][0] - x_mean)
	}

	// calculate the slope and intercept
	var slope, intercept float32
	slope = numerator / denominator
	intercept = y_mean - slope*x_mean

	return slope, intercept
}

func PredictY(x float32, slope float32, intercept float32) float32 {
	var y_pred = slope*x + intercept
	return y_pred
}

func CalculateRSquared(csv_object [][]float32, slope float32, intercept float32) float32 {
	var y_mean = MeanOfColumn(csv_object, 1)
	var numerator, denominator float32

	for i := 0; i < len(csv_object); i++ {
		var y_pred = PredictY(csv_object[i][0], slope, intercept)
		numerator += (y_pred - y_mean) * (y_pred - y_mean)
		denominator += (csv_object[i][1] - y_mean) * (csv_object[i][1] - y_mean)
	}

	var r_squared = numerator / denominator
	return r_squared
}
