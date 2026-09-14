package main

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

const testTolerance float32 = 1e-4

func assertFloat32Equal(t *testing.T, name string, got, want float32) {
	t.Helper()
	// Allow float32 rounding at the scale of the expected result, not just near zero.
	tolerance := float64(testTolerance) + 1e-6*math.Abs(float64(want))
	if math.IsNaN(float64(got)) || math.IsInf(float64(got), 0) || math.Abs(float64(got)-float64(want)) > tolerance {
		t.Fatalf("%s = %v, want %v", name, got, want)
	}
}

func TestReadCSV(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "data.csv")
	content := "1,2\n3.5,4.25\n"

	if err := os.WriteFile(filename, []byte(content), 0600); err != nil {
		t.Fatalf("failed to create test CSV: %v", err)
	}

	got, err := read_csv(filename)
	if err != nil {
		t.Fatalf("read_csv returned an unexpected error: %v", err)
	}

	want := [][]float32{{1, 2}, {3.5, 4.25}}
	if len(got) != len(want) {
		t.Fatalf("read_csv returned %d rows, want %d", len(got), len(want))
	}
	for i := range want {
		for j := range want[i] {
			assertFloat32Equal(t, "csv value", got[i][j], want[i][j])
		}
	}
}

func TestReadCSVMissingFile(t *testing.T) {
	_, err := read_csv(filepath.Join(t.TempDir(), "missing.csv"))
	if err == nil {
		t.Fatal("read_csv returned nil error for a missing file")
	}
}

func TestMeanOfCol(t *testing.T) {
	data := [][]float32{
		{1, 10},
		{2, 20},
		{3, 30},
	}

	assertFloat32Equal(t, "mean of column 0", mean_of_col(data, 0), 2)
	assertFloat32Equal(t, "mean of column 1", mean_of_col(data, 1), 20)
}

func TestCalculateSlopeAndIntercept(t *testing.T) {
	data := [][]float32{
		{1, 3},
		{2, 5},
		{3, 7},
		{4, 9},
	}

	slope, intercept := calculate_slope_and_intercept(data)
	assertFloat32Equal(t, "slope", slope, 2)
	assertFloat32Equal(t, "intercept", intercept, 1)
}

func TestPredictY(t *testing.T) {
	assertFloat32Equal(t, "predicted y", predict_y(4, 2, 1), 9)
}

func TestCalculateRSquared(t *testing.T) {
	data := [][]float32{
		{1, 3},
		{2, 5},
		{3, 7},
		{4, 9},
	}

	slope, intercept := calculate_slope_and_intercept(data)
	rSquared := calculate_r_squared(data, slope, intercept)
	assertFloat32Equal(t, "R-squared", rSquared, 1)
}

func TestLinearRegressionWithSampleData(t *testing.T) {
	data, err := read_csv("sample_data.csv")
	if err != nil {
		t.Fatalf("failed to read sample_data.csv: %v", err)
	}

	slope, intercept := calculate_slope_and_intercept(data)
	rSquared := calculate_r_squared(data, slope, intercept)

	assertFloat32Equal(t, "sample slope", slope, 0.30537275)
	assertFloat32Equal(t, "sample intercept", intercept, 156.90932)
	assertFloat32Equal(t, "sample R-squared", rSquared, 0.49953422)
	assertFloat32Equal(t, "sample prediction", predict_y(4512, slope, intercept), 1534.7512)
}
