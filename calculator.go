// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
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
	operators := []string{"*", "-", "+", "/"}
	//Check if it has spaces if so we let fmt.Sscanf do the token parsing for use if not then we split and get the operands and operator ourself.

	var operator string
	var operA, operB, result float64

	if len(strings.Fields(expr)) < 3 {
		//probably the expr did not have spaces where we expect them to be.
		for _, oper := range operators {
			if strings.Contains(expr, oper) {
				operator = oper
				break
			}
		}

		var err error

		x := strings.TrimSpace(strings.Split(expr, operator)[0])

		operA, err = strconv.ParseFloat(x, 64)
		if err != nil {
			return result, err
		}

		y := strings.TrimSpace(strings.Split(expr, operator)[1])

		operB, err = strconv.ParseFloat(y, 64)
		if err != nil {
			return result, err
		}
	} else {
		if _, err := fmt.Sscanf(expr, "%f%s%f", &operA, &operator, &operB); err != nil {
			return result, err
		}
	}

	//Now lets do the math here

	switch operator {
	case "*":
		result = operA * operB
	case "/":
		result = operA / operB
	case "-":
		result = operA - operB
	case "+":
		result = operA + operB
	default:
		return result, fmt.Errorf("calcstring: unsupported operator for input expression: '%v'", expr)
	}

	return result, nil
}
