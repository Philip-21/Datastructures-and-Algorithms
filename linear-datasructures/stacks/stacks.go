package main

import "fmt"

//Stacks are used in parsers for solving maze algorithms.
//stacks follows the List-in-first-out principle.(L-I-F-O)
//(last element added in a list will be the firt to remove)
//can illustrated using a slice in Go

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//push method adds items onto the stack
func (s *Stack) Push(str string) {
	//append the value onto the stack
	*s = append(*s, str)

}

//remove something from the stack
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		///get the index of the last element in the list
		index := len(*s) - 1   //gets the index of the top most element
		element := (*s)[index] //grab the value and storing it
		*s = (*s)[:index]      //removing the element (i.e slicing all the element but not including the element we want to remove )
		return element, true
	}
}

func main() {
	var stack Stack

	//add some values into the Stack
	stack.Push("Data Structures")
	stack.Push("And")
	stack.Push("Algorithms")

	///pop of all the values
	for !stack.IsEmpty() {
		x, y := stack.Pop() //returns 2 values from the pop method
		//if y is true we print the actual value x that we removed
		if y == true {
			fmt.Println(x)
		}
	}
	//when we run the program it prints out
	//(Algorithms and Data Structures) in a reverse format
	//which means algorithm is the last one that we added and the first one that gets removed

}
