package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"

    "gonum.org/v1/gonum/mat"
)

// ---------- Data Structures ----------
type DataPoint struct {
    X float64
    Y float64
}

// LinearModel wraps learned weights and exposes Predict methods.
type LinearModel struct {
    Weights []float64 // [intercept, slope]
}

// Predict returns the predicted value for a single input x.
func (m *LinearModel) Predict(x float64) float64 {
    y := m.Weights[0] + m.Weights[1]*x
    return y
}

// PredictBatch returns predictions for a slice of inputs.
func (m *LinearModel) PredictBatch(xs []float64) []float64 {
    preds := make([]float64, len(xs))
    for i, x := range xs {
        preds[i] = m.Predict(x)
    }
    return preds
}

// ---------- Utilities ----------
func readCSV(filename string) ([]DataPoint, error) {
    var dataset []DataPoint
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    r := csv.NewReader(f)
    for {
        record, err := r.Read()
        if err != nil {
            break
        }
        if len(record) < 2 {
            continue
        }
        x, errX := strconv.ParseFloat(record[0], 64)
        y, errY := strconv.ParseFloat(record[1], 64)
        if errX == nil && errY == nil {
            dataset = append(dataset, DataPoint{X: x, Y: y})
        }
    }
    return dataset, nil
}

// Train a linear regression model using Normal Equation
func trainLinearModel(dataset []DataPoint) *LinearModel {
    n := len(dataset)
    X := mat.NewDense(n, 2, nil)
    y := mat.NewVecDense(n, nil)

    for i, dp := range dataset {
        X.Set(i, 0, 1)     // intercept term
        X.Set(i, 1, dp.X)  // feature
        y.SetVec(i, dp.Y)
    }

    var xt mat.Dense
    xt.Mul(X.T(), X)

    var xtInv mat.Dense
    if err := xtInv.Inverse(&xt); err != nil {
        log.Fatal("Matrix not invertible:", err)
    }

    var xty mat.Dense
    xty.Mul(X.T(), y)

    var w mat.Dense
    w.Mul(&xtInv, &xty)

    weights := w.RawMatrix().Data
    return &LinearModel{Weights: weights}
}

// R² calculation
func rSquared(yTrue, yPred []float64) float64 {
    yMean := mean(yTrue)
    var ssRes, ssTot float64
    for i := range yTrue {
        ssRes += math.Pow(yTrue[i]-yPred[i], 2)
        ssTot += math.Pow(yTrue[i]-yMean, 2)
    }
    return 1 - ssRes/ssTot
}

func mean(data []float64) float64 {
    var sum float64
    for _, v := range data {
        sum += v
    }
    return sum / float64(len(data))
}

// ---------- Main ----------
func main() {
    // Example dataset file
    filepath := "sample_data.csv"

    dataset, err := readCSV(filepath)
    if err != nil {
        log.Fatal(err)
    }

    // Train model
    model := trainLinearModel(dataset)
    fmt.Printf("Intercept: %.4f, Slope: %.4f\n", model.Weights[0], model.Weights[1])

    // Predictions
    yTrue := make([]float64, len(dataset))
    for i, dp := range dataset {
        yTrue[i] = dp.Y
    }
    yPred := model.PredictBatch(yTrue) // predict using same Xs

    // Evaluate
    r2 := rSquared(yTrue, yPred)
    fmt.Printf("R²: %.4f\n", r2)

    // Predict new values
    fmt.Println("Predict(4512):", model.Predict(4512))
    fmt.Println("PredictBatch([3738, 4261, 3777]):", model.PredictBatch([]float64{3738, 4261, 3777}))
}
