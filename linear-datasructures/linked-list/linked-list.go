package main

import (
	"fmt"
)

//Linked List is a linear datastructure in which elements are not stored in contiguous memory locations

//arrays are good for modifying a certain element
//Linked List is good for inserting and deleting an elemen

//each element in a linked list is called a node
//a node is composed of two parts(1= Data, 2= Pointer or referece to the next node in the list)
//a node at the beginning of the linked list is referred to as HEAD

type node struct {
	data int   //Data
	next *node //pointer or reference

}

type linkedList struct {
	head   *node
	length int
}

// a prepend function that allows us to add node to the list
// it accepts a value and creates a new node that adds the data to the list
func (l *linkedList) Prepend(value int) {
	newnode := node{data: value}
	if l.head != nil {
		newnode.next = l.head //.next points to the one that was previously the first node in the list
		l.head = &newnode     //moving the head to the new node (next *node)
		l.length++            //increasin the length by 1
	} else {
		l.head = &newnode
		l.length++
	}

}

// a receiver funtion that prints out all the values stored in the list
func (l *linkedList) Printlist() bool {
	if l.head == nil {
		return false
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data) //printing out every value stored in the node
		currentNode = currentNode.next
		return true
	}
	return true
}

// a delete func to delete nodes from the list
func (l *linkedList) DeleteWithValue(value int) any {
	var novalue any
	var lengthReduced any
	var done any

	if l.length == 0 {
		return novalue
	}
	//using -----------{45,98,34,55 }-----as case study

	if l.head.data == value {
		l.head = l.head.next
		l.length-- //reduce the lenght of the list by 1
		return lengthReduced
		//returns 98,34,55
	}
	//choosing a particular node to delete

	//we are 1 node before the one we may have to delete

	//we want to delete 98
	PreviousToDelete := l.head //setting the node as {45,98,34,55}

	//instruction if the nxt node is not = to the value we want to delete
	for PreviousToDelete.next.data != value {
		//if we have hit the end of the list
		//it returns
		if PreviousToDelete.next.next == nil {
			return done
		}
		PreviousToDelete = PreviousToDelete.next
	}
	//span it to 2 nodes ahead i.e {45,34,55}
	PreviousToDelete.next = PreviousToDelete.next.next
	l.length-- //reduce the node by taking node 98 out
	return done
}

func main() {
	myList := linkedList{}
	myList.Prepend(21)
	myList.Prepend(31)
	myList.Prepend(56)
	myList.Prepend(65)
	myList.Prepend(86)
	myList.Printlist() //prints out 86, 65, 56,31, 21

	myList.DeleteWithValue(31)
	myList.Printlist() //prints out 86, 65, 56, 21
}
