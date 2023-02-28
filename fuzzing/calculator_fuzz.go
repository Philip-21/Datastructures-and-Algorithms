package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func Calculator(num1, num2 float64, operator string) (float64, error) {
	var result float64

	switch operator {
	case "-":
		result = num1 - num2

	case "+":
		result = num1 + num2

	case "/":
		if num2 == 0 || num1 == 0 {
			return 0, errors.New("Division by 0")
		}
	case "*":
		result = num1 * num2
	}

	return result, nil
}

func TestCalculator(t *testing.T) {
	testcases := []struct {
		num1     float64
		num2     float64
		operator string
		output   float64
	}{
		{1, 3, "+", 3},
		{4, 6, "-", -2},
		{0, 0, "/", 0},
		{9, 2, "*", 18},
	}
	for _, tc := range testcases {
		result, _ := Calculator(tc.num1, tc.num2, tc.operator)
		if result != tc.output {
			t.Errorf("Calculator(%d, %d, %s) = %s; expected %s", tc.num1, tc.num2, tc.operator,
				result, tc.output)

		}

	}
}

/*
The go-fuzz testing tools operates on raw binary data
Every input are in []bytes which are coverted to strings and
The string is then split into three parts (two numbers and an operator) using strings.
Split, and the numbers are parsed using strconv.ParseFloat. Finally,
the Calculator function is called with the parsed numbers and operator.
*/
func Fuzz(data []byte) float64 {
	//converts input data into string
	inputs := strings.Split(string(data), ",")
	//parse input numbers and operators
	num1, err := strconv.ParseFloat(inputs[0], 64)
	if err != nil {
		return 0

	}
	num2, err := strconv.ParseFloat(inputs[1], 64)
	if err != nil {
		return 0
	}
	operators := inputs[2]
	_, err = Calculator(num1, num2, operators)
	if err != nil {
		if err.Error() != "Division by 0" {
			return 0
		}
	}
	//returns a success
	return 1
}

func main() {
	// Seed the random number generator
	rand.Seed(1234)

	// Generate random input data
	data := make([]byte, 64)
	rand.Read(data)

	// Fuzz the Calculator function with the generated input data
	Fuzz(data)
}
