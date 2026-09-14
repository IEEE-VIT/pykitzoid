package metrics

import "testing"

func TestMeanSquaredError(t *testing.T) {
	got := MeanSquaredError([]float64{1, 2, 3}, []float64{2, 2, 5})
	if got != 4.0/3.0 {
		t.Fatalf("MeanSquaredError() = %v, want %v", got, 4.0/3.0)
	}
}

func TestMeanAbsoluteError(t *testing.T) {
	got := MeanAbsoluteError([]float64{1, 2, 3}, []float64{2, 2, 5})
	if got != 1.0 {
		t.Fatalf("MeanAbsoluteError() = %v, want 1", got)
	}
}

func TestMetricsEmptyInputs(t *testing.T) {
	if got := MeanSquaredError(nil, nil); got != 0 {
		t.Fatalf("MeanSquaredError() with empty inputs = %v, want 0", got)
	}
	if got := MeanAbsoluteError(nil, nil); got != 0 {
		t.Fatalf("MeanAbsoluteError() with empty inputs = %v, want 0", got)
	}
}

func TestMetricsRejectMismatchedInputs(t *testing.T) {
	for _, metric := range []struct {
		name string
		call func()
	}{
		{"MSE", func() { MeanSquaredError([]float64{1}, nil) }},
		{"MAE", func() { MeanAbsoluteError([]float64{1}, nil) }},
	} {
		t.Run(metric.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected mismatched inputs to panic")
				}
			}()
			metric.call()
		})
	}
}
