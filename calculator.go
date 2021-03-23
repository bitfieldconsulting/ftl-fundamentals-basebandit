// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes two or more numbers and returns the result of adding them together.
func Add(args ...float64) float64 {
	var result float64

	for _, arg := range args {
		result += arg
	}

	return result
}

// Subtract takes two or more numbers and returns the result of subtracting the second
// from the first.
func Subtract(args ...float64) float64 {
	var result float64

	for i, arg := range args {
		if i == 0 {
			result = arg
			continue
		}
		result -= arg
	}

	return result
}

//Multiply takes two or more numbers and returns the result of multiplying the second
// by the first.
func Multiply(args ...float64) float64 {
	var result float64 = 1

	for _, arg := range args {
		result *= arg
	}

	return result
}

//Divide takes two or more numbers and returns the result of dividing the second by the first.
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

//Sqrt takes a number and returns its square root
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, fmt.Errorf("sqrt: negative number '%f'", a)
	}
	return math.Sqrt(a), nil
}

//CalcString evaluates a string as a mathematical expression and yields the result otherwise returns error if encountered.
func CalcString(expr string) (float64, error) {
	operators := []string{"*", "-", "+", "/"}
	//Check if it has spaces if so we let fmt.Sscanf do the token parsing for use if not then we split and get the operands and operator ourself.

	var operator string
	var operA, operB, result float64

	fields := len(strings.Fields(expr))
	if fields < 3 {
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
	} else if fields > 3 {
		return result, fmt.Errorf("calcstring: does not support more that two operands %q", expr)
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
		return result, fmt.Errorf("calcstring: unsupported operator for input expression: %q", expr)
	}

	return result, nil
}
