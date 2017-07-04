package main

import "fmt"

type LinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	list := NewLinkedList()
	list.AddToFrontByNode(&Node{Value: 2})
	list.AddToFrontByValue(1)
	// list.printLinkedList()

	list.AddToEndByNode(&Node{Value: 3})
	list.AddToEndByValue(4)
	list.AddToEndByNode(&Node{Value: 5})
	// list.printLinkedList()

	list.RemoveFromEnd()

	fmt.Printf("4 is in list: %t\n", list.ContainsValue(4))
	list.RemoveByValue(4)
	fmt.Printf("4 is in list: %t\n", list.ContainsValue(4))

	list.AddToEndByValue(5)
	list.AddToEndByValue(5)
	list.AddToEndByValue(5)
	list.printLinkedList()
	list.RemoveByValue(5)
	list.printLinkedList()

	list.Clear()
	list.printLinkedList()

}

func NewNode(val int) *Node {
	return &Node{Value: val}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) AddToFrontByValue(val int) {
	n := NewNode(val)
	l.addToFrontInternal(n)
}

func (l *LinkedList) AddToFrontByNode(n *Node) {
	l.addToFrontInternal(n)
}

func (l *LinkedList) addToFrontInternal(n *Node) {
	if l.isEmpty() {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
	l.Count++
}

func (l *LinkedList) RemoveFromFront() {
	if l.isEmpty() {
		fmt.Println("Cannot remove from empty List")
		return
	}

	if l.Count == 1 {
		l.Head, l.Tail = nil, nil
	} else {
		l.Head = l.Head.Next
	}

	l.Count--
}

func (l *LinkedList) AddToEndByValue(val int) {
	n := NewNode(val)
	l.addToEndInternal(n)
}

func (l *LinkedList) AddToEndByNode(n *Node) {
	l.addToEndInternal(n)
}

func (l *LinkedList) addToEndInternal(n *Node) {
	if l.isEmpty() {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n

	l.Count++
}

func (l *LinkedList) RemoveFromEnd() {
	if l.isEmpty() {
		fmt.Println("Cannot remove from empty List")
		return
	}

	if l.Count == 1 {
		l.Head, l.Tail = nil, nil
	} else {
		current := l.Head
		for current.Next != l.Tail {
			current = current.Next
		}
		l.Tail, current.Next = current, nil
	}

	l.Count--
}

func (l *LinkedList) RemoveByValue(val int) {
	if l.isEmpty() {
		fmt.Println("Cannot remove by value from empty list")
	}
	if l.Head.Value == val {
		l.RemoveFromFront()
	}

	current := l.Head
	previous := current
	var found bool
	for current != nil {
		if current.Value == val {
			found = true
			if l.Tail == current {
				previous.Next = nil
			} else {
				previous.Next = current.Next
			}
			l.Count--

			l.printLinkedList()
		}
		previous = current
		current = current.Next
	}

	if found {
		fmt.Printf("Removed all instances of %d from the linked list\n", val)
	} else {
		fmt.Printf("%d is not a value in the linked list\n", val)
	}
}

func (l *LinkedList) isEmpty() bool {
	if l.Count == 0 {
		return true
	}
	return false
}

func (l *LinkedList) ContainsValue(val int) bool {
	if l.isEmpty() {
		return false
	}
	if l.Head.Value == val {
		return true
	}

	current := l.Head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}

	return false
}

func (l *LinkedList) Clear() {
	l.Head, l.Tail, l.Count = nil, nil, 0
}

func (l *LinkedList) printLinkedList() {
	fmt.Println("Starting to print list")
	if l.isEmpty() {
		fmt.Println("Empty List")
		return
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}
