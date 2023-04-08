package main

import (
	"fmt"
)

func ForLoop() {
	//for loops are use to repeat instructions for a particular code

	//The simplest counter-based iteration, the basic form is:
	//for initial statement; conditional statement; modified statement {}

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

	//exmaple 4
	//Give the odd and even numbers of particluar items
	for z := 0; z < 5; z++ {
		fmt.Println(z)
		if z%2 == 0 { //odd number that will give remainders willpass through this statement
			z = z / 2
		} else {
			z = 2*z + z //for even nmbers
		}
	}

}

func Range() {
	//Range iterates over elements

	//Repeat the number of elements btwn 1-10
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range n {
		fmt.Println(i)
	}

	// for range loop
	sv := []int{1, 2, 3, 4, 5}
	for item := range sv {
		fmt.Println(item)
	}

}

func NestedLoop() {
loop:
	for s := 1; s <= 3; s++ {
		for q := 3; q <= 7; q++ {
			fmt.Println(q * s)
			if q*s >= 4 {
				break loop
			}
		}
	}
}

func If() {
	//if statements are condition statements

	//example 1
	//display the even numbers between 1-20 using if statements

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var s int
	for q := 1; q <= len(num); q = q + 2 {
		if s == q {
			fmt.Printf("these are the even numbers %d/n:", num[q])
		}
	}

	//example 2
	//Make a program that divides x by 2 if it’s greater than 0
	var x = 18
	var d = 2
	if x > 0 {
		v := x / d
		fmt.Println(v)
	}

	///example 3
	//write an if statement for conditioning the length of strings
	first := []string{"Hrr", "the", "trol", "trav"}
	second := []string{"th", "li", "ty", "ol", "pi"}
	third := []string{"te", "oi", "ju", "or", "te", "teo"}
	if len(first) > len(second) {
		fmt.Println("wrong statement")
	} else {
		if len(second) == len(third) {
			fmt.Println("Wrong Statement")
		} else {
			if len(third) > len(first) && len(second) < len(third) {
				fmt.Println("Right Statement")
			}
		}
	}

}

func Switch() {
	//this is a shorter way to 	write a sequence of if else statement

	//example 1
	id := 50
	switch {
	case id < 40:
		fmt.Println("less than 40")
	case id <= 50:
		fmt.Println("fifty")
		fallthrough //this enables the nextcase to be executed
	case id > 40:
		fmt.Println("greater than forty")

	}

}

func main() {
	ForLoop()
	Range()
	NestedLoop()
	If()
	Switch()

}
