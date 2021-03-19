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
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 8
	got := calculator.Multiply(4.5, 2)
	if want != got {
		t.Errorf("want %f,got %f", want, got)
	}
}
