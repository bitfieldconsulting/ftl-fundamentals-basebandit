package calculator_test

import (
	"calculator"
	"testing"
)

//testCase contains a set of inputs and expected outputs for each test run.
type testCase struct {
	inputs []float64
	output float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{inputs: []float64{3.0, 2.0}, output: 5.0},
		{inputs: []float64{0.50, 0.2}, output: 0.70},
		{inputs: []float64{200, 30}, output: 230},
		{inputs: []float64{4.56, 3.56}, output: 8.12},
	}

	for _, tt := range tests {
		got := calculator.Add(tt.inputs[0], tt.inputs[1])
		if tt.output != got {
			t.Errorf("want %f, got %f", tt.output, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{inputs: []float64{-3.0, 2.0}, output: -5.0},
		{inputs: []float64{-0.50, -0.2}, output: -0.3},
		{inputs: []float64{200, 30}, output: 170},
		{inputs: []float64{0.56, -3.56}, output: 4.12},
		{inputs: []float64{0.0, -0.0}, output: 0.0},
		{inputs: []float64{0, 0}, output: -1},
	}

	for _, tt := range tests {
		got := calculator.Subtract(tt.inputs[0], tt.inputs[1])
		if tt.output != got {
			t.Errorf("want %f, got %f", tt.output, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{inputs: []float64{-3.0, 2.0}, output: -6.0},
		{inputs: []float64{-0.50, -0.2}, output: 0.1},
		{inputs: []float64{200, 30}, output: 6000},
		{inputs: []float64{0.56, -3.56}, output: -1.9936000000000003},
		{inputs: []float64{0.0, -0.0}, output: -0.0},
		{inputs: []float64{0, 0}, output: 1},
	}

	for _, tt := range tests {
		got := calculator.Multiply(tt.inputs[0], tt.inputs[1])
		if tt.output != got {
			t.Errorf("want %f,got %f", tt.output, got)
		}
	}
}
