package main

import "fmt"

func main() {

	arr := []int{9, 6, 7, 5, 20}
	LB := 0 //lower bound

	// the index of the last element is n-1. n is the number of element
	//so 5-1 = 4 last index is 4
	//Therefore, len(arr) gives us the total number of elements in the array,
	// not the index of the last element.
	UB := len(arr) - 1 //Upper bound which is the last element of the list

	//Initialize the counter
	k := LB

	//traverse
	for k <= UB {
		//visit each element in the array
		fmt.Println(arr[k])
		//increase the counter
		k++

	}

	// Print the total number of elements in the array
	fmt.Println("Total number of elements in the array:", UB-LB+1)
}
