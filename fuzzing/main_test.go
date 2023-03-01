package main

import (
	"errors"
	"fmt"
	"testing"
)

func calculator(num1, num2 float64, operator string) (float64, error) {
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

func fuzzCalculator(data float64) float64 {
	inputcases := []struct {
		input1   float64
		input2   float64
		operator string
		output   float64
	}{
		{1, 3, "+", 4},
		{4, 6, "-", -2},
		{0, 0, "/", 0},
		{9, 2, "*", 18},
		{1.2, 2.4, "*", 2.88},
		{2, 6, "+", 9},
	}
	success := 1
	for _, input := range inputcases {
		result, _ := calculator(input.input1, input.input2, input.operator)
		if result != input.output {
			fmt.Errorf("Calculator(%f,%f,%s)=%f; expected %f", input.input1, input.input2,
				input.operator, result, input.output)
		}
		return result
	}
	return float64(success)

}

func FuzzTest(f *testing.F) {
	f.Fuzz(func(t *testing.T, data float64) {
		_ = fuzzCalculator(data)
	})

}

func TestCalculator(t *testing.T) {
	testcases := []struct {
		num1     float64
		num2     float64
		operator string
		output   float64
	}{
		{1, 3, "+", 4},
		{4, 6, "-", -2},
		{0, 0, "/", 0},
		{9, 2, "*", 18},
	}
	for _, tc := range testcases {
		result, _ := calculator(tc.num1, tc.num2, tc.operator)
		if result != tc.output {
			t.Errorf("Calculator(%f,%f,%s)=%f; expected %f", tc.num1, tc.num2, tc.operator,
				result, tc.output)

		}

	}
}
