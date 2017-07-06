package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	linkList *LinkedList
}

type LinkedList struct {
	Head  *Node
	Count int
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	queue := &Queue{&LinkedList{}}

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
	if q.linkList.Count == 0 {
		q.linkList.Head = n
	} else {
		current := q.linkList.Head
		for current != nil {
			current = current.Next
		}
		current.Next = n
	}
	q.linkList.Count++
}

func (q *Queue) dequeue() (int, error) {
	if q.linkList.Count == 0 {
		return 0, errors.New("Cannot dequeue value from empty queue")
	}

	tmp := q.linkList.Head.Value
	q.linkList.Head = q.linkList.Head.Next
	q.linkList.Count--
	return tmp, nil
}

func (q *Queue) peek() (int, error) {
	if q.linkList.Count == 0 {
		return 0, errors.New("Cannot peek in to empty queue")
	}
	return q.linkList.Head.Value, nil
}
