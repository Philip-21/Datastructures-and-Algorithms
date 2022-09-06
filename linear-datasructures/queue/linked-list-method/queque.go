package main

import (
	"container/list"
	"fmt"
)

// an alternative way of creating a queque is using a linked List
//is using the list package package in Go

func main() {
	queque := list.New()

	//PushBack inserts a new element e
	//with value v at the back of list l and returns e
	queque.PushBack("Data Strutures")
	queque.PushBack("And")
	queque.PushBack("Algorithms")

	//Front1 returns the first element of list
	front1 := queque.Front()
	fmt.Println(front1.Value) //The value stored with this element.
	queque.Remove(front1)

	//front 2 returns the second element
	front2 := queque.Front()
	fmt.Println(front2.Value) //The value stored with this element.
	queque.Remove(front2)

	//front 3 returns the third element
	front3 := queque.Front()
	fmt.Println(front3.Value) //The value stored with this element.
	queque.Remove(front3)
}
