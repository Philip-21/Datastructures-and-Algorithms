package main

//queue consists of elements to be processed in a particular order or based on priority
//Queues are commonly used for storing tasks that
//need to be done, or incoming HTTP requests that need to be processed by a server
//handling interruptions in real-time systems, call handling, and CPU task scheduling
//are good examples for using queues

// Queue—Array of OrdersType
type Queue []*Order

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
func (que *Queue) Add(ord *Order) {
	if len(*que) == 0 {

		*que = append(*que, ord)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *que {
			if ord.cost > addedOrder.cost {
				*que = append((*que)[:i], append(Queue{ord}, (*que)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*que = append(*que, ord)
		}
	}

}
