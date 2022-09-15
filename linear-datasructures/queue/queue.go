package main

import "fmt"

//queue consists of elements to be processed in a particular order or based on priority
//Queues are commonly used for storing tasks that
//need to be done, or incoming HTTP requests that need to be processed by a server
//handling interruptions in real-time systems, call handling, and CPU task scheduling
//are good examples for using queues

//Queue is an ordered data structure which follows the FIFO first-in-first-out principle
//can be illustrated using a slice or a linked list

///--------Example1
type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// adds items onto the queque
func (q *Queue) Enqueque(str string) {
	//append the value onto the queque
	*q = append(*q, str)

}

//remove something from the queque
func (q *Queue) Dequeque() (string, bool) {
	if q.IsEmpty() {
		return "", false
	} else {
		///get the index of the first element in the list
		element := (*q)[0] //grab the value and storing it
		*q = (*q)[1:]      // slicing all the elements from the first index
		return element, true
	}
}

func Example1() {
	var que Queue
	que.Enqueque("Data Structures")
	que.Enqueque("And")
	que.Enqueque("Algorithms")

	///pop of all the values
	for !que.IsEmpty() {
		x, y := que.Dequeque() //returns 2 values from the Dequeque method
		//if y is true we print the actual value x that we removed
		if y == true {
			fmt.Println(x)
		}
	}
	//prints  data Structures and algorithms
}

//Example 2
// Queues—Array of OrdersType
type Queues []*Order

//The Order class
type Order struct {
	cost         int
	quantity     int
	product      string
	customerName string
}

// New method initializes the propeties of the ord with cost, quantity, product,customerName
func (ord *Order) New(cost int, quantity int, product string, customerName string) {
	ord.cost = cost
	ord.quantity = quantity
	ord.product = product
	ord.customerName = customerName
}

//the Add method on the Queue class takes the ord
//parameter and adds it to Queue based on the cost
func (que *Queues) Add(ord *Order) {
	if len(*que) == 0 {

		*que = append(*que, ord)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *que {
			if ord.cost > addedOrder.cost {
				*que = append((*que)[:i], append(Queues{ord}, (*que)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*que = append(*que, ord)
		}
	}

}

//

func main() {
	Example1()
}
