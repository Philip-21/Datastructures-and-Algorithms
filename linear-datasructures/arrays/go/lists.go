package main

import (
	"container/list"
	"fmt"
)

//data structure refers to the organization of data in a computer's memory, in order to retrieve it quickly for processing.

func main() {
	items := [8]int{1, 12, 34, 45, 55, 67, 55, 32} ///7 indexins
	th := items[2:7]                               //gets the 2nd element to the 6th elemnt
	fmt.Println(items)
	fmt.Println(th)

	//(2)SLICES
	sl := [8]int{2, 3, 4, 55, 66, 7, 12, 34}
	fmt.Printf("length : %v\n", len(sl))
	fmt.Printf("capacity : %v\n", cap(sl))
	cl := []int{1, 2, 3, 4, 5, 76, 23, 89, 56, 40, 100}
	cb := cl[:]   //slice all elements
	df := cl[3:]  // slice from 4th element
	db := cl[2:8] // slice from the 3rd element to 8th element
	fmt.Println(cl)
	fmt.Println(cb)
	fmt.Println(df)
	fmt.Println(db)

	// applying append () in slicing
	cv := []int{}
	fmt.Println(cv)
	fmt.Printf("length: %v\n", len(cv))
	fmt.Printf("capacity: %v\n", cap(cv))
	cv1 := append(cv, []int{2, 3, 4, 5}...)
	fmt.Println(cv1)
	fmt.Printf("length: %v\n", len(cv1))
	//stack operations in slicing
	stack := []int{61, 982, 763, 246, 985}
	stack1 := stack[:len(stack)-1] //getting the first 4 elements
	fmt.Println(stack1)

	var intlist list.List
	intlist.PushBack(13) //pushback is used to insert list node or value  at the end of a list
	intlist.PushBack(21)
	intlist.PushBack(39)
	//     for element := intlist.Front() ; element != nil; element=element.Next()
	// 	fmt.Println(element.Value.(int))

}

//the list is iterated through the for loop,
// and the element's value is accessed through the Value method.
