package main

// import whatever packages you will require here
import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// use the following sample dataset
/*
| hours_studied | passed |
|----------------|--------|
| 0.50           | 0      |
| 1.75           | 0      |
| 1.75           | 1      |
| 3.25           | 1      |
| 5.00           | 1      |
*/

// expect the dataset to be in csv form, where the first column is the
// feature (x) and the second column is the binary class label (0 or 1)

// FUNCTION DEFINITION

// function to read a csv and return the dataset as [][]float64
func read_csv(filename string) ([][]float64, error) {
	// Read the file into a string
	// if error occurs, return an empty slice and the error
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a new reader to evaluate the file as a csv
	// if error occurs, return an empty slice and the error
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	// convert the string array to float array
	var lines_converted [][]float64
	for i := 0; i < len(lines); i++ {
		x, err := strconv.ParseFloat(lines[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseFloat(lines[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}
		lines_converted = append(lines_converted, []float64{x, y})
	}
	return lines_converted, nil
}

// sigmoid squashes any real-valued number into the range (0, 1) so it
// can be interpreted as a probability

// // FORMULA
// // sigmoid(z) = 1 / (1 + e^(-z))

func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// predict_probability returns the model's predicted probability that
// the given x belongs to class 1
func predict_probability(x float64, weight float64, bias float64) float64 {
	return sigmoid(weight*x + bias)
}

// predict_class converts a predicted probability into a binary class
// label using a 0.5 decision threshold
func predict_class(probability float64) int {
	if probability >= 0.5 {
		return 1
	}
	return 0
}

// compute_cost calculates the binary cross-entropy (log) loss of the
// model over the dataset

// // FORMULA
// // J(w,b) = -(1/n) * Σ[y*log(p) + (1-y)*log(1-p)]
// // where p is the predicted probability and y is the true label

func compute_cost(csv_object [][]float64, weight float64, bias float64) float64 {
	// clamp predictions away from 0 and 1 so log() never receives 0
	const epsilon = 1e-15

	var cost float64
	n := float64(len(csv_object))

	for i := 0; i < len(csv_object); i++ {
		x, y := csv_object[i][0], csv_object[i][1]
		p := predict_probability(x, weight, bias)
		p = math.Min(math.Max(p, epsilon), 1-epsilon)
		cost += y*math.Log(p) + (1-y)*math.Log(1-p)
	}

	return -cost / n
}

// gradient_descent fits the weight and bias by minimising the binary
// cross-entropy loss over the given number of epochs, and returns the
// cost history so training progress can be inspected or plotted

// // FORMULA
// // dw = (1/n) * Σ (p_i - y_i) * x_i
// // db = (1/n) * Σ (p_i - y_i)
// // weight -= learning_rate * dw
// // bias   -= learning_rate * db

func gradient_descent(csv_object [][]float64, learning_rate float64, epochs int) (float64, float64, []float64) {
	var weight, bias float64
	n := float64(len(csv_object))
	cost_history := make([]float64, 0, epochs)

	for epoch := 0; epoch < epochs; epoch++ {
		var dw, db float64
		for i := 0; i < len(csv_object); i++ {
			x, y := csv_object[i][0], csv_object[i][1]
			p := predict_probability(x, weight, bias)
			dw += (p - y) * x
			db += p - y
		}

		weight -= learning_rate * (dw / n)
		bias -= learning_rate * (db / n)

		cost_history = append(cost_history, compute_cost(csv_object, weight, bias))
	}

	return weight, bias, cost_history
}

// calculate_accuracy measures the fraction of samples the model
// classifies correctly
func calculate_accuracy(csv_object [][]float64, weight float64, bias float64) float64 {
	var correct float64
	for i := 0; i < len(csv_object); i++ {
		x, y := csv_object[i][0], csv_object[i][1]
		predicted := predict_class(predict_probability(x, weight, bias))
		if float64(predicted) == y {
			correct++
		}
	}
	return correct / float64(len(csv_object))
}

// function to plot the data points, showing their true class on the y-axis

// plot_data_points creates a scatter plot of the data and saves it as a PNG file.
// csv_object: The data points to be plotted.
// plot_name: The name of the PNG file to save the plot.
func plot_data_points(csv_object [][]float64, plot_name string) {
	// Create a new plot
	p := plot.New()

	// Customize the appearance
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)
	p.Title.Text = "Logistic Regression - Data Points"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Class"

	// Prepare the data points to be plotted
	pts := make(plotter.XYs, len(csv_object))
	for i := range csv_object {
		pts[i].X = csv_object[i][0]
		pts[i].Y = csv_object[i][1]
	}

	// Create a scatter plotter and add the points to it
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}

	p.Add(scatter)

	// Save the plot to the specified PNG file
	err = p.Save(1500, 900, plot_name)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Data points plotted successfully and saved as ", plot_name)
}

// function to plot the fitted sigmoid curve against the data points

func plot_sigmoid_curve(csv_object [][]float64, weight float64, bias float64, plot_name string) {
	p := plot.New()

	p.Title.Text = "Logistic Regression - Fitted Sigmoid Curve"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Predicted Probability"

	// find the range of x values so the curve is drawn across the data
	x_min, x_max := csv_object[0][0], csv_object[0][0]
	for i := range csv_object {
		if csv_object[i][0] < x_min {
			x_min = csv_object[i][0]
		}
		if csv_object[i][0] > x_max {
			x_max = csv_object[i][0]
		}
	}

	const steps = 100
	curve := make(plotter.XYs, steps+1)
	for i := 0; i <= steps; i++ {
		x := x_min + (x_max-x_min)*float64(i)/steps
		curve[i].X = x
		curve[i].Y = predict_probability(x, weight, bias)
	}

	line, err := plotter.NewLine(curve)
	if err != nil {
		log.Panic(err)
	}
	p.Add(line)

	points := make(plotter.XYs, len(csv_object))
	for i := range csv_object {
		points[i].X = csv_object[i][0]
		points[i].Y = csv_object[i][1]
	}
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Panic(err)
	}
	p.Add(scatter)

	err = p.Save(1500, 900, plot_name)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Sigmoid curve plotted successfully and saved as ", plot_name)
}

// MAIN FUNCTION

func main() {

	// define the csv file path
	var filepath string
	filepath = "sample_data.csv"

	// read the csv file into a file reader object
	var csv_object, err = read_csv(filepath)

	// if error occurs, print the error and exit
	if err != nil {
		log.Fatal(err)
	}

	// hyperparameters for gradient descent
	const learning_rate = 0.1
	const epochs = 5000

	// fit the model using gradient descent on the binary cross-entropy loss
	weight, bias, _ := gradient_descent(csv_object, learning_rate, epochs)

	// evaluate the fitted model
	cost := compute_cost(csv_object, weight, bias)
	accuracy := calculate_accuracy(csv_object, weight, bias)

	// print the weight, bias, cost and accuracy
	log.Println("weight: ", weight)
	log.Println("bias: ", bias)
	log.Println("final cost: ", cost)
	log.Println("accuracy: ", accuracy)

	// plot the data points and the fitted sigmoid curve
	plot_data_points(csv_object, "DataPoints.png")
	plot_sigmoid_curve(csv_object, weight, bias, "SigmoidCurve.png")
}
