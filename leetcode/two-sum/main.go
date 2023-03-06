package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target.
*/
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	//creating a hash map to store numbers
	/*
			the hash map makes this algorithm possible by allowing constant-time lookups of
			 the complement of each element in the nums slice.
			 The algorithm iterates over each element in the nums slice and checks if the complement of the current element exists in the hash map.
			  If the complement exists in the hash map, the function has found a pair of elements that add up to the target,
		      and the function returns the indices of these two elements
	*/
	/*
		THIS IS THE BIG O NOTATION ALGORITHM
		Using a hash map to store the elements of the nums slice allows the function to achieve a time complexity of O(n) because
		the lookup and insertion operations into the hash map take constant time.
		 This is much faster than the O(n^2) time complexity of a brute-force algorithm that compares each pair of elements in the nums slice.
	*/
	m := make(map[int]int)

	/*
		 iterates over each element in the nums slice and calculates
		 the complement of the current element that would add up to the target
		The complement is then checked if it exists in the hash map. If the complement exists,
		 it means that the function has found a pair of elements that add up to the target
	*/
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
			//adds te current index to the hash map
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
