package main

import "fmt"

func ForLoop() {
	//The simplest counter-based iteration, the basic form is:
	//for initial statement; conditional statement; modified statement {}
	//for loops are use to repeat instructions for a particular code

	//example 1
	//list the number of items from 0-10
	for i := 0; i < 11; i++ {
		fmt.Printf("The items from 0-10  are : %d\n", i)
	}

	//example 2
	//list the number of item from 1-20 incrementing by 2
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 0; i < len(a); i += 2 { //i<len(a) means i < the number of elements in the array of a
		fmt.Printf("items increasing by 2 are : %d\n", a[i])
		//this wil print out odd numbers
		fmt.Printf("odd Numbers between 1-20 are: %d\n", a[i])
	}

	///example 3
	//list the even number from the above
	for v := 1; v <= len(a); v = v + 2 {
		fmt.Println("Even Numbers btwn 1-20 are:", a[v])
	}

}

func Range() {
	//Range iterates over elements

	//Repeat the number of elements btwn 1-10
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range n {
		fmt.Println(i)
	}

}

func main() {
	ForLoop()
	Range()
}
