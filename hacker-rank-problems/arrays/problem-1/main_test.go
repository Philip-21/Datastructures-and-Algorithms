package main

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 4, 5, 6}, []int{6, 5, 4, 1}},
		{[]int{7, 8, 9, 10}, []int{10, 9, 8, 7}},
		{[]int{3, 4, 5, 9}, []int{9, 5, 4, 3}},
	}

	for _, c := range cases {
		result, _ := reverseArray(c.input)
		//package reflect to compare results
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("reverseArray(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}
