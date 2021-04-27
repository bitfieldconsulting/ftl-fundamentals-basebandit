// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strings"
)

// Add takes atleast two or more numbers and returns the result of adding them together.
func Add(a, b float64, args ...float64) float64 {

	if len(args) > 0 {
		sum := a + b

		for _, arg := range args {
			sum += arg
		}

		return sum
	}

	return a + b
}

// Subtract takes two and or possibly more numbers and returns the result of subtracting consequently from the results
// of the first arguments.
func Subtract(a, b float64, args ...float64) float64 {

	if len(args) > 0 {
		result := a - b
		for _, arg := range args {
			result -= arg
		}
		return result
	}

	return a - b
}

// Multiply takes two or more numbers and returns the result of multiplying the second
// by the first.
func Multiply(args ...float64) float64 {
	var result float64 = 1

	for _, arg := range args {
		result *= arg
	}

	return result
}

// Divide takes two or more numbers and returns the result of dividing the second by the first.
func Divide(args ...float64) (float64, error) {
	var result float64

	for i, arg := range args {
		if i > 0 && arg == 0 { //the first argument can be zero no problem but subsequent args cannot be: why? divide by zero runtime error
			return 0, fmt.Errorf("divide: divide by zero for input %v", args[i])
		}

		if i == 0 { //for the first iteration.
			result = arg
			continue
		}

		result = result / arg
	}

	return result, nil
}

// Sqrt takes a number and returns its square root
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, fmt.Errorf("sqrt: negative number '%f'", a)
	}
	return math.Sqrt(a), nil
}

// CalcString evaluates a string as a mathematical expression and yields the result otherwise returns error if encountered. A valid expression should have spaces in between the opreands and the operator.
func CalcString(expr string) (float64, error) {
	//we let fmt.Sscanf do the token parsing for us if the number of fields in the epxression is 3; that is 2 operands and 1 operator.

	var operator string
	var operA, operB, result float64

	fields := len(strings.Fields(expr))

	if fields > 3 {
		return result, fmt.Errorf("calcstring: does not support more than two operands %q", expr)
	}

	if _, err := fmt.Sscanf(expr, "%f%s%f", &operA, &operator, &operB); err != nil {
		return result, err
	}
	fmt.Println(operA, operator, operB)
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
		return result, fmt.Errorf("calcstring: unsupported operator for input expression: %q", expr)
	}

	return result, nil
}
