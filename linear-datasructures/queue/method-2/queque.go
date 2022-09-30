package main

//creating a Queue program using the O(n) notation
type Node struct {
	Value int
}

type Queue []*Node

func (q *Queue) Pop(n *Node) int {
	n = (*q)[0]   //getting the first element
	*q = (*q)[1:] ///takes out the first element
	return n.Value
}

func (q *Queue) Push(n *Node) {
	*q = append(*q, n)
}

func main() {
	var que Queue

	que.Push(&Node{1})
	que.Push(&Node{2})
	que.Push(&Node{3})

}
