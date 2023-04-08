package main

import "fmt"

func main() {
	//traversing
	v := []int{12, 34, 55, 67}
	for _, val := range v {
		println(val)
	}
	val := []int{34, 65, 86, 90, 21}
	next := append(val, 56)
	for i := 0; i < len(val); i++ {
		fmt.Println("Real value", val)
		fmt.Println("temporary value", next)
	}
}
