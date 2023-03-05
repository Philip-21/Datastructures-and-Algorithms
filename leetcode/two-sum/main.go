package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target.
*/
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	//creating a hash map to store numbers
	m := make(map[int]int, len(nums))

	for idx, val := range nums {
		i, ok := m[target-val]
		if ok {
			/*
				If the if statement condition is true,
				this line returns a slice of integers that contains the indices i and idx,
				 which represent the indices of the two
				 elements whose sum is equal to the input target
			*/
			fmt.Println([]int{i, idx})
			return []int{i, idx}
		} else {
			m[val] = idx
		}

	}
	return result

}

func main() {
	tc := []int{2, 7, 11, 15}
	target := 9

	_ = twoSum(tc, target)
	//prints a the indices[0,1]
	//which means 2,7 add up to give 9

}
