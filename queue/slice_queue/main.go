package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Nodes []Node
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	nodes := make([]Node, 0)
	queue := &Queue{nodes}

	queue.queue(&Node{Value: 1})

	if peek, err := queue.peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Peeking found %d\n", peek)
	}

	if val, err := queue.dequeue(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Dequeued %d\n", val)
	}

	if peek, err := queue.peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Peeking found %d\n", peek)
	}

	if val, err := queue.dequeue(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Dequeued %d\n", val)
	}
}

func (q *Queue) queue(n *Node) {
	q.Nodes = append([]Node{*n}, q.Nodes...)
}

func (q *Queue) dequeue() (int, error) {
	if len(q.Nodes) == 0 {
		return 0, errors.New("Cannot dequeue value from empty queue")
	}

	tmp := q.Nodes[0].Value
	q.Nodes = q.Nodes[1:]

	return tmp, nil
}

func (q *Queue) peek() (int, error) {
	if len(q.Nodes) == 0 {
		return 0, errors.New("Cannot peek in to empty queue")
	}

	return q.Nodes[0].Value, nil
}
