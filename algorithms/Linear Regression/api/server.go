// Package api exposes versioned HTTP endpoints for Pykitzoid algorithms.
package api

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"mime"
	"net/http"

	"github.com/IEEE-VIT/pykitzoid/algorithms/Linear_Regression/regression"
)

const MaxBodyBytes = 1 << 20
const MaxSamples = 10000

// Bound accumulations in the repository's float32 regression implementation.
const MaxInputMagnitude = 1e15

type sample struct {
	X *float32 `json:"x"`
	Y *float32 `json:"y"`
}

type fitRequest struct {
	Samples []sample   `json:"samples"`
	Predict []*float32 `json:"predict"`
}

type fitResponse struct {
	Algorithm   string    `json:"algorithm"`
	Slope       float32   `json:"slope"`
	Intercept   float32   `json:"intercept"`
	RSquared    *float32  `json:"r_squared"`
	SampleCount int       `json:"sample_count"`
	Predictions []float32 `json:"predictions"`
}

// NewHandler returns an isolated router; new algorithms can register here.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if requireMethod(w, r, http.MethodGet) {
			writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
		}
	})
	mux.HandleFunc("/api/v1/algorithms", func(w http.ResponseWriter, r *http.Request) {
		if requireMethod(w, r, http.MethodGet) {
			writeJSON(w, http.StatusOK, map[string]any{"algorithms": []map[string]string{
				{"id": "linear-regression", "method": "POST", "endpoint": "/api/v1/linear-regression", "description": "Fit y = slope*x + intercept and predict values."},
			}})
		}
	})
	mux.HandleFunc("/api/v1/linear-regression", fitLinearRegression)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusNotFound, "Endpoint not found.")
	})
	return mux
}

func requireMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method == method {
		return true
	}
	w.Header().Set("Allow", method)
	writeError(w, http.StatusMethodNotAllowed, "Method not allowed.")
	return false
}

func fitLinearRegression(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		writeError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json.")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var request fitRequest
	if err := decoder.Decode(&request); err != nil {
		decodeError(w, err)
		return
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		decodeError(w, err)
		return
	}
	if len(request.Samples) < 2 || len(request.Samples) > MaxSamples {
		writeError(w, http.StatusBadRequest, "Provide between 2 and 10000 samples.")
		return
	}
	if len(request.Predict) > MaxSamples {
		writeError(w, http.StatusBadRequest, "Provide at most 10000 prediction inputs.")
		return
	}
	data := make([][]float32, len(request.Samples))
	for i, point := range request.Samples {
		if point.X == nil || point.Y == nil || !finite(*point.X) || !finite(*point.Y) {
			writeError(w, http.StatusBadRequest, "Every sample must have finite numeric x and y values.")
			return
		}
		data[i] = []float32{*point.X, *point.Y}
		if math.Abs(float64(*point.X)) > MaxInputMagnitude || math.Abs(float64(*point.Y)) > MaxInputMagnitude {
			writeError(w, http.StatusUnprocessableEntity, "Sample magnitudes must not exceed 1e15; rescale the data.")
			return
		}
	}
	var variesX, variesY bool
	for _, point := range data[1:] {
		variesX = variesX || point[0] != data[0][0]
		variesY = variesY || point[1] != data[0][1]
	}
	if !variesX {
		writeError(w, http.StatusBadRequest, "Samples must contain at least two distinct x values.")
		return
	}
	for _, x := range request.Predict {
		if x == nil || !finite(*x) {
			writeError(w, http.StatusBadRequest, "Prediction inputs must be finite numbers.")
			return
		}
	}
	slope, intercept := regression.CalculateSlopeAndIntercept(data)
	if !finite(slope) || !finite(intercept) {
		writeError(w, http.StatusUnprocessableEntity, "Values exceed the algorithm's numeric range; rescale the data.")
		return
	}
	result := fitResponse{
		Algorithm: "linear-regression", Slope: slope, Intercept: intercept,
		SampleCount: len(data), Predictions: make([]float32, len(request.Predict)),
	}
	// R-squared is undefined when all observed y values are identical.
	if variesY {
		score := regression.CalculateRSquared(data, slope, intercept)
		if !finite(score) {
			writeError(w, http.StatusUnprocessableEntity, "Cannot compute R-squared at this numeric scale; rescale the data.")
			return
		}
		result.RSquared = &score
	}
	for i, x := range request.Predict {
		prediction := regression.PredictY(*x, slope, intercept)
		if !finite(prediction) {
			writeError(w, http.StatusUnprocessableEntity, "Prediction exceeds the algorithm's numeric range.")
			return
		}
		result.Predictions[i] = prediction
	}
	writeJSON(w, http.StatusOK, result)
}

func finite(value float32) bool {
	return !math.IsNaN(float64(value)) && !math.IsInf(float64(value), 0)
}

func decodeError(w http.ResponseWriter, err error) {
	var tooLarge *http.MaxBytesError
	if errors.As(err, &tooLarge) {
		writeError(w, http.StatusRequestEntityTooLarge, "Request body exceeds 1 MiB.")
		return
	}
	writeError(w, http.StatusBadRequest, "Provide one valid JSON object with numeric samples and prediction inputs, and no unknown fields.")
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
