package main

import "fmt"

//Create an array with the number 0 to 10
func Exercise1() {
	new := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int
	//A for loop is used for accessing all the elements in an array
	for i = 0; i < len(new); i++ {
		fmt.Println("array of numbers", new[i])
	}
}

func Exercise2() {
	names := []string{"Dennis", "Harry", "Ben", "Mike", "Colwil"}
	fmt.Println(names)
}

func main() {
	Exercise1()
	Exercise2()
}
