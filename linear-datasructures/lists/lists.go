package main

import (
	"container/list"
	"fmt"
)

//data structure refers to the organization of data in a computer's memory, in order to retrieve it quickly for processing.

func main(){
	var intlist list.List
	intlist.PushBack(13)//pushback is used to insert list node or value  at the end of a list 
	intlist.PushBack(21)
	intlist.PushBack(39)
    for element := intlist.Front() ; element != nil; element=element.Next()
	fmt.Println(element.Value.(int))

}

//the list is iterated through the for loop,
// and the element's value is accessed through the Value method.