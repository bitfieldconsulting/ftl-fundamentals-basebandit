package calculator_test

import (
	"calculator"
	"testing"
)

//testCase contains a set of inputs and expected outputs for each test run.
type testCase struct {
	name   string
	inputs []float64
	output float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{name: "Two positive fractional numbers which sum to positive", inputs: []float64{3.0, 2.0}, output: 5.0},
		{name: "Two positive fractional numbers which sum to positive", inputs: []float64{0.50, 0.2}, output: 0.70},
		{name: "Two positive numbers which sum to positive", inputs: []float64{200, 30}, output: 230},
		{name: "Two positive fractional numbers which sum to positive", inputs: []float64{4.56, 3.56}, output: 8.12},
		{name: "Two negative numbers which sum to positive", inputs: []float64{-2.0, -3.0}, output: -5.0},
		{name: "One positive and one negative number which sum to negative", inputs: []float64{-5.0, 2.0}, output: -3.0},
	}

	for _, tt := range tests {
		got := calculator.Add(tt.inputs[0], tt.inputs[1])
		if tt.output != got {
			t.Errorf("%s:\n\t\t\twant %f, got %f", tt.name, tt.output, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{name: "One negative number and one positive number which subtract to negative", inputs: []float64{-3.0, 2.0}, output: -5.0},
		{name: "Two negative numbers which subtract to negative", inputs: []float64{-0.50, -0.2}, output: -0.3},
		{name: "Two positive numbers which subtract to positive", inputs: []float64{200, 30}, output: 170},
		{name: "One positive number and one negative number which subtract to positive", inputs: []float64{0.56, -3.56}, output: 4.12},
		{name: "One positive number and one negative number which subtract to negative", inputs: []float64{0.0, -0.0}, output: 0.0},
		{inputs: []float64{0, 0}, output: -1},
	}

	for _, tt := range tests {
		got := calculator.Subtract(tt.inputs[0], tt.inputs[1])
		if tt.output != got {
			t.Errorf("%s:\n\t\t\twant %f, got %f", tt.name, tt.output, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{name: "One negative number and one positive number which multiply to negative", inputs: []float64{-3.0, 2.0}, output: -6.0},
		{name: "Two negative numbers which multiply to positive", inputs: []float64{-0.50, -0.2}, output: 0.1},
		{name: "Two positive numbers which multiply to positive", inputs: []float64{200, 30}, output: 6000},
		{name: "One positive number and one negative number which multiply to negative", inputs: []float64{0.56, -3.56}, output: -1.9936000000000003},
		{name: "One positive number and one negative number which multiply to negative", inputs: []float64{0.0, -0.0}, output: -0.0},
		{inputs: []float64{0, 0}, output: 1},
	}

	for _, tt := range tests {
		got := calculator.Multiply(tt.inputs[0], tt.inputs[1])
		if tt.output != got {
			t.Errorf("%s:\n\t\t\twant %f,got %f", tt.name, tt.output, got)
		}
	}
}
