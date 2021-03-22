// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b, c float64) float64 {
	return a + b + c
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b, c float64) float64 {
	return (a - b) - c
}

//Multiply takes two numbers and returns the result of multiplying the second
// by the first.
func Multiply(a, b, c float64) float64 {
	return (a * b) * c
}

//Divide takes two numbers and returns the result of dividing the second by the first.
func Divide(a, b, c float64) (float64, error) {
	if b == 0 || c == 0 {
		return 0.0, errors.New("divide: divide by zero")
	}
	return (a / b) / c, nil
}

//Sqrt takes a number and returns its square root
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, errors.New("sqrt: negative numbers don't have real square roots")
	}
	return math.Sqrt(a), nil
}
