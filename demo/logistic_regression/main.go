// Command logistic_regression_demo demonstrates the Logistic Regression
// algorithm (see algorithms/Logistic Regression) end-to-end: it trains a
// binary classifier on sample_data.csv using gradient descent over the
// binary cross-entropy loss, then uses the fitted model to classify a
// few new, unseen inputs.
//
// Run it with:
//
//	go run .
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type dataPoint struct {
	x float64
	y float64
}

func readCSV(filename string) ([]dataPoint, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	data := make([]dataPoint, 0, len(rows))
	for _, row := range rows {
		x, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}
		data = append(data, dataPoint{x: x, y: y})
	}
	return data, nil
}

// sigmoid(z) = 1 / (1 + e^(-z))
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

func predictProbability(x, weight, bias float64) float64 {
	return sigmoid(weight*x + bias)
}

func predictClass(probability float64) int {
	if probability >= 0.5 {
		return 1
	}
	return 0
}

// fit trains weight and bias with gradient descent on the binary
// cross-entropy loss:
//
//	dw = (1/n) * Σ (p_i - y_i) * x_i
//	db = (1/n) * Σ (p_i - y_i)
func fit(data []dataPoint, learningRate float64, epochs int) (weight, bias float64) {
	n := float64(len(data))

	for epoch := 0; epoch < epochs; epoch++ {
		var dw, db float64
		for _, point := range data {
			p := predictProbability(point.x, weight, bias)
			dw += (p - point.y) * point.x
			db += p - point.y
		}
		weight -= learningRate * (dw / n)
		bias -= learningRate * (db / n)
	}

	return weight, bias
}

func accuracy(data []dataPoint, weight, bias float64) float64 {
	var correct float64
	for _, point := range data {
		predicted := predictClass(predictProbability(point.x, weight, bias))
		if float64(predicted) == point.y {
			correct++
		}
	}
	return correct / float64(len(data))
}

func main() {
	data, err := readCSV("sample_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	const learningRate = 0.1
	const epochs = 5000

	weight, bias := fit(data, learningRate, epochs)

	fmt.Printf("Trained model: weight=%.4f, bias=%.4f\n", weight, bias)
	fmt.Printf("Training accuracy: %.2f%%\n", accuracy(data, weight, bias)*100)

	fmt.Println("\nPredicting on new, unseen inputs:")
	for _, hoursStudied := range []float64{1.5, 2.8, 4.2, 6.0} {
		probability := predictProbability(hoursStudied, weight, bias)
		class := predictClass(probability)
		outcome := "Fail"
		if class == 1 {
			outcome = "Pass"
		}
		fmt.Printf("  hours_studied=%.1f -> P(pass)=%.4f -> %s\n", hoursStudied, probability, outcome)
	}
}
