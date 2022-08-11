package main

import "strconv"

//Stacks are used in parsers for solving maze algorithms.

//Element class
type Element struct {
	elementValue int
}

//string method on Element class
func (ele *Element) StringMethod() string {
	return strconv.Itoa(ele.elementValue) //StringMethod   returns the elementvalue
}

//The Stack class has the count and array pointer of elements
type Stack struct {
	elements     []*Element
	elementCount int
}

//The NewMethod on the Stack class creates a dynamic array of elements.

func (sta *Stack) New() {
	sta.elements= make(*Element[] elements,0)
}

