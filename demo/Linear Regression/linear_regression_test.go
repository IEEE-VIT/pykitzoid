func TestLinearRegression(t *testing.T) {
    // Example test case for linear regression
    result := LinearRegression([]float64{1, 2, 3}, []float64{1, 2, 3})
    expected := 1.0
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}