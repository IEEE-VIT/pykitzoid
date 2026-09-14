package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func send(method, path, contentType, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	r.Header.Set("Content-Type", contentType)
	w := httptest.NewRecorder()
	NewHandler().ServeHTTP(w, r)
	return w
}

const validRequest = `{"samples":[{"x":1,"y":3},{"x":2,"y":5},{"x":3,"y":7}],"predict":[4,5]}`

func TestFitAndPredict(t *testing.T) {
	w := send("POST", "/api/v1/linear-regression", "application/json", validRequest)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	var result fitResponse
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Slope != 2 || result.Intercept != 1 || result.RSquared == nil || *result.RSquared != 1 || result.SampleCount != 3 {
		t.Fatalf("unexpected fit: %+v", result)
	}
	if len(result.Predictions) != 2 || result.Predictions[0] != 9 || result.Predictions[1] != 11 {
		t.Fatalf("unexpected predictions: %v", result.Predictions)
	}
}

func TestSampleCSV(t *testing.T) {
	file, err := os.Open("../sample_data.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	points := make([]map[string]float32, len(rows))
	for i, row := range rows {
		x, err := strconv.ParseFloat(row[0], 32)
		if err != nil {
			t.Fatal(err)
		}
		y, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			t.Fatal(err)
		}
		points[i] = map[string]float32{"x": float32(x), "y": float32(y)}
	}
	body, _ := json.Marshal(map[string]any{"samples": points, "predict": []float32{4512}})
	w := send("POST", "/api/v1/linear-regression", "application/json", string(body))
	var result fitResponse
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil || w.Code != 200 {
		t.Fatalf("response: %s, error: %v", w.Body, err)
	}
	if result.RSquared == nil || math.Abs(float64(*result.RSquared)-0.49953422) > 1e-4 || math.Abs(float64(result.Predictions[0])-1534.7512) > 1e-3 {
		t.Fatalf("unexpected sample result: %+v", result)
	}
}

func TestConstantY(t *testing.T) {
	w := send("POST", "/api/v1/linear-regression", "application/json", `{"samples":[{"x":1,"y":5},{"x":2,"y":5}]}`)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"r_squared":null`) || !strings.Contains(w.Body.String(), `"predictions":[]`) {
		t.Fatalf("unexpected result: %d %s", w.Code, w.Body)
	}
}

func TestInvalidRequests(t *testing.T) {
	cases := []struct {
		name, body string
		status     int
	}{
		{"empty body", "", 400},
		{"broken JSON", "{", 400},
		{"null", "null", 400},
		{"empty samples", `{"samples":[]}`, 400},
		{"too few samples", `{"samples":[{"x":1,"y":2}]}`, 400},
		{"constant x", `{"samples":[{"x":1,"y":2},{"x":1,"y":3}]}`, 400},
		{"missing x", `{"samples":[{"y":2},{"x":2,"y":3}]}`, 400},
		{"null y", `{"samples":[{"x":1,"y":null},{"x":2,"y":3}]}`, 400},
		{"string x", `{"samples":[{"x":"1","y":2},{"x":2,"y":3}]}`, 400},
		{"unknown field", strings.TrimSuffix(validRequest, "}") + `,"typo":1}`, 400},
		{"extra JSON", validRequest + `{}`, 400},
		{"null prediction", `{"samples":[{"x":1,"y":2},{"x":2,"y":3}],"predict":[null]}`, 400},
		{"out of range number", `{"samples":[{"x":1e100,"y":2},{"x":2,"y":3}]}`, 400},
		{"arithmetic overflow", `{"samples":[{"x":1e30,"y":2},{"x":2e30,"y":3}]}`, 422},
		{"too many samples", `{"samples":[` + strings.Repeat(`{"x":1,"y":2},`, MaxSamples) + `{"x":2,"y":3}]}`, 400},
		{"too many predictions", strings.Replace(validRequest, `[4,5]`, `[`+strings.Repeat("1,", MaxSamples)+`1]`, 1), 400},
		{"oversized body", strings.Repeat(" ", MaxBodyBytes+1), 413},
		{"oversized trailing data", validRequest + strings.Repeat(" ", MaxBodyBytes), 413},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := send("POST", "/api/v1/linear-regression", "application/json", c.body)
			if w.Code != c.status {
				t.Fatalf("wanted %d, got %d: %s", c.status, w.Code, w.Body)
			}
			var result map[string]string
			if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil || result["error"] == "" {
				t.Fatalf("missing JSON error: %s", w.Body)
			}
		})
	}
}

func TestRouting(t *testing.T) {
	for _, path := range []string{"/health", "/api/v1/algorithms"} {
		if w := send("GET", path, "", ""); w.Code != 200 || !json.Valid(w.Body.Bytes()) {
			t.Fatalf("GET %s: %d %s", path, w.Code, w.Body)
		}
		if w := send("POST", path, "", ""); w.Code != 405 || w.Header().Get("Allow") != "GET" {
			t.Fatalf("POST %s: %d", path, w.Code)
		}
	}
	if w := send("GET", "/api/v1/linear-regression", "", ""); w.Code != 405 || w.Header().Get("Allow") != "POST" {
		t.Fatalf("expected method error: %d", w.Code)
	}
	if w := send("POST", "/api/v1/linear-regression", "text/plain", validRequest); w.Code != 415 {
		t.Fatalf("expected content type error: %d", w.Code)
	}
	if w := send("GET", "/missing", "", ""); w.Code != 404 {
		t.Fatalf("expected not found: %d", w.Code)
	}
}

func TestConcurrentRequests(t *testing.T) {
	handler := NewHandler()
	var group sync.WaitGroup
	for i := 0; i < 20; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/api/v1/linear-regression", strings.NewReader(validRequest))
			r.Header.Set("Content-Type", "application/json; charset=utf-8")
			handler.ServeHTTP(w, r)
			if w.Code != 200 {
				t.Error(fmt.Sprintf("concurrent request failed: %d", w.Code))
			}
		}()
	}
	group.Wait()
}
