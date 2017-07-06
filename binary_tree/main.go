package main

import "fmt"

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	binaryTree := &BinaryTree{}
	binaryTree.Add(&Node{Value: 4})
	fmt.Printf("binaryTree: %+v\n", binaryTree.Root)

	binaryTree.Add(&Node{Value: 3})
	fmt.Printf("binaryTree: %+v\n", binaryTree.Root.Left)

	binaryTree.Add(&Node{Value: 5})
	fmt.Printf("binaryTree: %+v\n", binaryTree.Root.Right)
	fmt.Printf("binaryTree: %+v\n", binaryTree.Root)

	if val, err := binaryTree.Find(6); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found: %d\n", val)
	}

}

func (b *BinaryTree) Add(n *Node) {
	if b.Root == nil {
		b.Root = n
	} else {
		b.addInternal(b.Root, n)
	}
}

func (b *BinaryTree) addInternal(currentNode *Node, n *Node) {
	if n.Value < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = n
		} else {
			b.addInternal(currentNode.Left, n)
		}
	} else {
		if currentNode.Right == nil {
			currentNode.Right = n
		} else {
			b.addInternal(currentNode.Right, n)
		}
	}
}

func (b *BinaryTree) Remove(n *Node) {
	if b.Root == nil {
		b.Root = n
	} else {
		b.addInternal(b.Root, n)
	}
}

func (b *BinaryTree) Find(value int) (int, error) {
	if b.Root == nil {
		return 0, fmt.Errorf("unable to find %d in BinaryTree", value)
	}
	return b.findInternal(b.Root, value)
}

func (b *BinaryTree) findInternal(currentNode *Node, value int) (int, error) {
	if currentNode.Value == value {
		return currentNode.Value, nil
	}

	if value < currentNode.Value {
		if currentNode.Left == nil {
			return 0, fmt.Errorf("unable to find %d in BinaryTree", value)
		}
		return b.findInternal(currentNode.Left, value)
	}
	if currentNode.Right == nil {
		return 0, fmt.Errorf("unable to find %d in BinaryTree", value)
	}
	return b.findInternal(currentNode.Right, value)
}
