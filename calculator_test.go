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
		{name: "Two positive fractional numbers which sum to positive", inputs: []float64{3.0, 2.0, 1.0}, output: 6.0},
		{name: "Two positive fractional numbers which sum to positive", inputs: []float64{0.50, 0.2, 0.3}, output: 1.0},
		{name: "Two positive numbers which sum to positive", inputs: []float64{200, 30, 50}, output: 280},
		{name: "Three positive fractional numbers which sum to positive", inputs: []float64{4.56, 3.56, 1.06}, output: 9.18},
		{name: "Two negative fractional numbers which sum to positive", inputs: []float64{-2.0, -3.0, -2.0}, output: -7.0},
		{name: "One negative and two positive numbers which sum to negative", inputs: []float64{-5.0, 2.0, 0.0}, output: -3.0},
	}

	for _, tt := range tests {
		got := calculator.Add(tt.inputs[0], tt.inputs[1], tt.inputs[2])
		if tt.output != got {
			t.Errorf("%s:\n\t\t\twant %f, got %f", tt.name, tt.output, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{name: "One negative fractional number and two positive fractional numbers which subtract to negative", inputs: []float64{-3.0, 2.0, 0.1}, output: -5.1},
		{name: "Three negative fractional numbers which subtract to negative", inputs: []float64{-0.50, -0.2, -0.1}, output: -0.19999999999999998},
		{name: "Three positive numbers which subtract to positive", inputs: []float64{200, 30, 50}, output: 120},
		{name: "One positive fractional number and two negative fractional numbers which subtract to positive", inputs: []float64{0.56, -3.56, -0.2}, output: 4.32},
		{name: "One positive fractional number and two fractional negative numbers which subtract to negative", inputs: []float64{0.0, -0.0, -0.0}, output: 0.0},
		{inputs: []float64{0, 0, 0}, output: 0},
	}

	for _, tt := range tests {
		got := calculator.Subtract(tt.inputs[0], tt.inputs[1], tt.inputs[2])
		if tt.output != got {
			t.Errorf("%s:\n\t\t\twant %f, got %f", tt.name, tt.output, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{name: "Two negative fractional number and one positive fractional number which multiply to negative", inputs: []float64{-3.0, 2.0, -1.0}, output: 6.0},
		{name: "Two negative fractional numbers and one positive fractional number which multiply to positive", inputs: []float64{-0.50, -0.2, 0.1}, output: 0.010000000000000002},
		{name: "Three positive numbers which multiply to positive", inputs: []float64{200, 30, 10}, output: 60000},
		{name: "One positive fractional number and two negative fractional numbers which multiply to negative", inputs: []float64{0.56, -3.56, -0.1}, output: 0.19936000000000004},
		{name: "One positive fractional number and one negative fractional number which multiply to negative", inputs: []float64{0.0, -0.0, -0.0}, output: 0.0},
		{inputs: []float64{0, 0, 0}, output: 0},
	}

	for _, tt := range tests {
		got := calculator.Multiply(tt.inputs[0], tt.inputs[1], tt.inputs[2])
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
		{name: "One positive number and two negative numbers which divide to negative", inputs: []float64{1, -1, -1}, output: 1.0, errExpected: false},
		{name: "Three positive numbers which divide to positive", inputs: []float64{4, 2, 1}, output: 2, errExpected: false},
		{name: "Three positive numbers which do not divide", inputs: []float64{6, 0, 0}, errExpected: true},
		{name: "Three positive numbers which do not divide", inputs: []float64{6, 0, 1}, errExpected: true},
		{name: "Two negative numbers which divide to negative", inputs: []float64{-3, -5, -2}, output: -0.3, errExpected: false},
	}

	for _, tt := range tests {
		got, err := calculator.Divide(tt.inputs[0], tt.inputs[1], tt.inputs[2])
		if tt.errExpected {
			if err == nil {
				t.Errorf("%s:\n\t\t\twant error got %v", tt.name, err)
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
		c := rand.Float64()
		want := a + b + c
		got := calculator.Add(a, b, c)
		if want != got {
			t.Errorf("want %f got %f", want, got)
		}
	}
}

func TestSqrt(t *testing.T) {

	tests := []struct {
		name        string
		inputs      []float64
		output      float64
		errExpected bool
	}{
		{name: "A positive number whose square root is a real number", inputs: []float64{64}, output: 8, errExpected: false},
		{name: "A negative number whose square root is not a real number", inputs: []float64{-64}, errExpected: true},
		{name: "A positive number whose sqaure root is a real number", inputs: []float64{0}, output: 0, errExpected: false},
		{name: "A fractional positive number whose square root is a real number", inputs: []float64{0.64}, output: 0.8, errExpected: false},
	}

	for _, tt := range tests {
		got, err := calculator.Sqrt(tt.inputs[0])
		if tt.errExpected {
			if err == nil {
				t.Errorf("%s:\n\t\t\twant error got %v", tt.name, err)
			}
		}

		if tt.output != got {
			t.Errorf("%s:\n\t\t\twant %f got %f", tt.name, tt.output, got)
		}
	}
}

func TestCalcString(t *testing.T) {
	tests := []struct {
		input       string
		output      float64
		errExpected bool
	}{
		{input: "2*2", output: 4, errExpected: false},
		{input: "1+1.5", output: 2.5, errExpected: false},
		{input: "100-0.1", output: 99.9, errExpected: false},
		{input: "100-  0.1", output: 99.9, errExpected: false},
		{input: "18  /  6", output: 3, errExpected: false},
		{input: "99&1", errExpected: true},
		{input: "7 * 1 * 3", errExpected: true},
	}

	for _, tt := range tests {
		got, err := calculator.CalcString(tt.input)

		if err == nil && tt.errExpected { //we expected an non-nil error but got a nil error
			t.Errorf("want non-nil error got nil error %q", err)
		} else if (err != nil) && !tt.errExpected { //we did not expect an error but we got one anyway
			t.Errorf("want nil error got non-nil error %q", err)
		} else { //else our logic is wrong.
			if tt.output != got {
				t.Errorf("want %f got %f", tt.output, got)
			}
		}
	}
}
