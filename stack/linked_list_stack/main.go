package main

import (
	"errors"
	"fmt"
)

// Pros of linked list over slice
// 1. No hard size limit.
// 2. Easy to implement. No out of bounds checks.

// Cons of linked list over slice
// 1. Adding to stack causes memory allocation on every push.
// 2. Memory cost of each node can be more than cost of the data to store.
// 3. Can lower performation. They store data all over the heap

type Stack struct {
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
	s := &Stack{&LinkedList{}}
	s.Push(1)
	if a, err := s.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Peeking found %d at the top of the stack\n", a)
	}
	s.Push(2)
	if a, err := s.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Peeking found %d at the top of the stack\n", a)
	}
	s.Push(3)
	if a, err := s.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Peeking found %d at the top of the stack\n", a)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Poped: %d\n", v)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Poped: %d\n", v)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Poped: %d\n", v)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Poped: %d\n", v)
	}
}

func (s *Stack) Push(val int) {
	if s.linkList.Count == 0 {
		s.linkList.Head = &Node{Value: val}
	} else {
		s.linkList.Head = &Node{val, s.linkList.Head}
	}
	s.linkList.Count++
}

func (s *Stack) Pop() (int, error) {
	if s.linkList.Count == 0 {
		return 0, errors.New("Cannot pop from empty stack")
	}

	tmp := s.linkList.Head.Value
	s.linkList.Head = s.linkList.Head.Next
	s.linkList.Count--
	return tmp, nil
}

func (s *Stack) Peek() (int, error) {
	if s.linkList.Count == 0 {
		return 0, errors.New("Nothing to peek at. Stack is empty!")
	}
	return s.linkList.Head.Value, nil
}
