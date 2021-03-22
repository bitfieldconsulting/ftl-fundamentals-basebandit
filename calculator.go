// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
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

func CalcString(expr string) (float64, error) {

	operators := map[string]string{"multiplication": "*", "division": "/", "addition": "+", "subtraction": "-"}

	var operator string

	for k, v := range operators {
		if strings.Contains(strings.TrimSpace(expr), v) {
			operator = k
			break
		}
	}

	x := strings.Split(expr, operators[operator])[0]
	operA, err := strconv.ParseFloat(x, 64)

	if err != nil {
		panic(err)
	}

	y := strings.Split(expr, operators[operator])[1]
	operB, err := strconv.ParseFloat(y, 64)

	if err != nil {
		panic(err)
	}

	var result float64

	switch operator {
	case "multiplication":
		result = operA * operB
	case "division":
		result = operA / operB
	case "subtraction":
		result = operA - operB
	case "addition":
		result = operA + operB

	default:
		return result, errors.New("calcstring: unknown operand")
	}

	return result, nil
}
