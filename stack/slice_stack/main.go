package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Nodes []Node
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	nodes := make([]Node, 0)
	s := &Stack{Nodes: nodes}

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
		fmt.Printf("Popped: %d\n", v)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped: %d\n", v)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped: %d\n", v)
	}

	if v, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped: %d\n", v)
	}
}

func (s *Stack) Push(val int) {
	n := &Node{Value: val}
	s.Nodes = append(s.Nodes, *n)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Nodes) == 0 {
		return 0, errors.New("Cannot pop from empty stack")
	}
	tmp := s.Nodes[len(s.Nodes)-1].Value
	s.Nodes = s.Nodes[:len(s.Nodes)-1]

	return tmp, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.Nodes) == 0 {
		return 0, errors.New("Nothing to peek at. Stack is empty!")
	}

	return s.Nodes[len(s.Nodes)-1].Value, nil
}
