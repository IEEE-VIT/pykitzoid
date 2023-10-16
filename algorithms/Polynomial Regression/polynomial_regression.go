package main

// import whatever packages you will require here
import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

// use the following sample dataset
/*
| x     | y     |
|-------|-------|
| 4512  | 1530  |
| 3738  | 1297  |
| 4261  | 1335  |
| 3777  | 1282  |
| 4177  | 1590  |
*/

// expect the dataset to be in csv form

// FUNCTION DEFINITION
// add parameters and change the return type as necessary

// function to read csv file (dataset)
type DataPoint struct {
	X float64
	Y float64
}

func read_csv(filename string) ([]DataPoint, error) {
	var dataset []DataPoint

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		x, errX := strconv.ParseFloat(record[0], 64)
		y, errY := strconv.ParseFloat(record[1], 64)
		if errX == nil && errY == nil {
			dataPoint := DataPoint{X: x, Y: y}
			dataset = append(dataset, dataPoint)
		}
	}

	return dataset, nil
}

// function to calculate weights

// // FORMULA
// W = Inverse(X' . X) . X' . y
// where X' is the transpose of X and . denotes dot product

func calc_weights() {

}

// function to predict y values

// // FORMULA
// y = X . W
// Where W is the calculated weights and . denotes dot product

func predict_y(dataset []DataPoint, weights []float64) []float64 {
	var predictions []float64

	for _, dataPoint := range dataset {
		x := dataPoint.X
		y := weights[0]
		for i := 1; i < len(weights); i++ {
			y += weights[i] * x
			x *= dataPoint.X
		}
		predictions = append(predictions, y)
	}

	return predictions
}

// function to calculate mean
func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	var sum float64 = 0

	for _, value := range data {
		sum += value
	}

	return sum / float64(len(data))
}

// function to plot regression line
func plot_regression_line(dataset []DataPoint, weights []float64) error {
	p, err := plot.New()
	plotter.NewScatter(dataPoints)
	if err != nil {
		return err
	}

	points := make(plotter.XYs, len(dataset))
	for i, dp := range dataset {
		points[i].X = dp.X
		points[i].Y = dp.Y
	}

	s, err := plotter.NewScatter(points)
	if err != nil {
		return err
	}
	p.Add(s)

	// Generate points for the regression line
	xmin := dataset[0].X
	xmax := dataset[len(dataset)-1].X
	step := (xmax - xmin) / 100 // 100 points for the regression line
	regressionLine := make(plotter.XYs, 100)
	for i := range regressionLine {
		x := xmin + float64(i)*step
		y := weights[0]
		for j := 1; j < len(weights); j++ {
			y += weights[j] * x
			x *= x // for polynomial regression
		}
		regressionLine[i].X = x
		regressionLine[i].Y = y
	}
	l, err := plotter.NewLine(regressionLine)
	if err != nil {
		return err
	}
	p.Add(l)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.png"); err != nil {
		return err
	}

	return nil
}

// function to plot the data points

func plot_data_points() {
	// insert code here
}

// function to calculate R-squared value to measure accuracy of model

// // FORMULA
// R^2 = (Σ(y_pred - y')^2)/(Σ(y-y')^2)
// where y_pred is the predicted value for y, y' is the mean

func calculate_r_squared(csv_object [][]float32, slope float32, intercept float32) float32 {

	var xData []float64
	var yData []float64
	for i := 0; i < len(csv_object); i++ {
		xData = append(xData, float64(csv_object[i][0]))
		yData = append(yData, float64(csv_object[i][1]))
	}

	// Calculate the mean of the observed values (yData)
	yMean := mean(yData)

	var numerator, denominator float32

	for i := 0; i < len(csv_object); i++ {
		y_pred := predict_y(float32(xData[i]), slope, intercept)
		numerator += (y_pred - float32(yData[i])) * (y_pred - float32(yData[i]))
		denominator += (csv_object[i][1] - yMean) * (csv_object[i][1] - yMean)
	}

	r_squared := numerator / denominator
	return r_squared
}

// MAIN FUNCTION

func main() {

}
