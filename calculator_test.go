package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
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
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		inputs      []float64
		output      float64
		errExpected bool
	}{
		{name: "One positive number and one negative number which divide to negative", inputs: []float64{1, -1}, output: -1.0, errExpected: false},
		{name: "Two positive numbers which divide to positive", inputs: []float64{4, 2}, output: 3, errExpected: false},
		{name: "Two positive numbers which do not divide", inputs: []float64{6, 0}, errExpected: true},

		{name: "Two negative numbers which divide to negative", inputs: []float64{-3, -5}, output: 0.6, errExpected: false},
	}

	for _, tt := range tests {
		got, err := calculator.Divide(tt.inputs[0], tt.inputs[1])
		if tt.errExpected { //we expect an error
			if err != nil { //therefore err must be non-nil(runtime error)
				t.Errorf("%s: %v\n", tt.name, err)
			}
		}

		if got != tt.output { //we don't expect an error but we the output might be differenc from what we expect with a non nil error.
			t.Errorf("%s:\n\t\t\twant %f got %f", tt.name, tt.output, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("want %f got %f", want, got)
		}
	}
}
