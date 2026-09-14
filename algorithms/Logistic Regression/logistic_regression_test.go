package main

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

const testTolerance = 1e-6

func assertFloatEqual(t *testing.T, name string, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > testTolerance {
		t.Fatalf("%s = %v, want %v", name, got, want)
	}
}

func TestReadCSV(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "data.csv")
	content := "1,0\n3.5,1\n"

	if err := os.WriteFile(filename, []byte(content), 0600); err != nil {
		t.Fatalf("failed to create test CSV: %v", err)
	}

	got, err := read_csv(filename)
	if err != nil {
		t.Fatalf("read_csv returned an unexpected error: %v", err)
	}

	want := [][]float64{{1, 0}, {3.5, 1}}
	if len(got) != len(want) {
		t.Fatalf("read_csv returned %d rows, want %d", len(got), len(want))
	}
	for i := range want {
		for j := range want[i] {
			assertFloatEqual(t, "csv value", got[i][j], want[i][j])
		}
	}
}

func TestReadCSVMissingFile(t *testing.T) {
	_, err := read_csv(filepath.Join(t.TempDir(), "missing.csv"))
	if err == nil {
		t.Fatal("read_csv returned nil error for a missing file")
	}
}

func TestSigmoid(t *testing.T) {
	assertFloatEqual(t, "sigmoid(0)", sigmoid(0), 0.5)

	if got := sigmoid(20); got <= 0.999999 {
		t.Fatalf("sigmoid(20) = %v, want a value very close to 1", got)
	}
	if got := sigmoid(-20); got >= 0.000001 {
		t.Fatalf("sigmoid(-20) = %v, want a value very close to 0", got)
	}
}

func TestPredictProbability(t *testing.T) {
	assertFloatEqual(t, "predicted probability", predict_probability(0, 1, 0), 0.5)
}

func TestPredictClass(t *testing.T) {
	if predict_class(0.5) != 1 {
		t.Fatal("predict_class(0.5) should classify as 1 at the decision boundary")
	}
	if predict_class(0.49) != 0 {
		t.Fatal("predict_class(0.49) should classify as 0")
	}
	if predict_class(0.51) != 1 {
		t.Fatal("predict_class(0.51) should classify as 1")
	}
}

func TestComputeCost(t *testing.T) {
	// a confident, correct model should have cost close to 0
	confident := [][]float64{{10, 1}, {-10, 0}}
	cost := compute_cost(confident, 1, 0)
	if cost > 1e-3 {
		t.Fatalf("expected near-zero cost for a confident correct model, got %v", cost)
	}

	// an unbiased model (weight=0, bias=0) predicts p=0.5 for every
	// sample, which always costs exactly -log(0.5)
	uncertain := [][]float64{{1, 0}, {1, 1}}
	cost = compute_cost(uncertain, 0, 0)
	assertFloatEqual(t, "uncertain cost", cost, -math.Log(0.5))
}

func TestGradientDescentSeparatesLinearlySeparableData(t *testing.T) {
	data := [][]float64{
		{-3, 0}, {-2, 0}, {-1, 0},
		{1, 1}, {2, 1}, {3, 1},
	}

	weight, _, costHistory := gradient_descent(data, 0.5, 2000)

	accuracy := calculate_accuracy(data, weight, 0)
	if accuracy != 1.0 {
		t.Fatalf("expected perfect accuracy on linearly separable data, got %v", accuracy)
	}
	if weight <= 0 {
		t.Fatalf("expected a positive weight since larger x implies class 1, got %v", weight)
	}
	if len(costHistory) != 2000 {
		t.Fatalf("expected 2000 cost history entries, got %d", len(costHistory))
	}
	if costHistory[len(costHistory)-1] >= costHistory[0] {
		t.Fatalf("expected cost to decrease over training, started at %v ended at %v", costHistory[0], costHistory[len(costHistory)-1])
	}
}

func TestCalculateAccuracy(t *testing.T) {
	data := [][]float64{{1, 1}, {-1, 0}, {5, 1}, {-5, 0}}
	// weight=1, bias=0 classifies negative x as 0 and positive x as 1
	accuracy := calculate_accuracy(data, 1, 0)
	assertFloatEqual(t, "accuracy", accuracy, 1.0)
}

func TestLogisticRegressionWithSampleData(t *testing.T) {
	data, err := read_csv("sample_data.csv")
	if err != nil {
		t.Fatalf("failed to read sample_data.csv: %v", err)
	}

	weight, bias, _ := gradient_descent(data, 0.1, 5000)
	accuracy := calculate_accuracy(data, weight, bias)

	if weight <= 0 {
		t.Fatalf("expected a positive weight since more hours studied implies a higher pass probability, got %v", weight)
	}
	if accuracy < 0.75 {
		t.Fatalf("expected at least 0.75 accuracy on the sample dataset, got %v", accuracy)
	}
}
