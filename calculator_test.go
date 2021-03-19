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

	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
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
