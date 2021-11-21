package main

import (
	"fmt"
)

type Node struct {
	id          int
	left, right *Node
}

func NewNode(id int) *Node {
	return &Node{id: id}
}

func (n *Node) Invert() {
	if n.left != nil {
		n.left.Invert()
	}
	if n.right != nil {
		n.right.Invert()
	}
	n.left, n.right = func() (*Node, *Node) { return n.right, n.left }()
}
func (n Node) String() string {
	var l, r string
	if n.left != nil {
		l = n.left.String()
	} else {
		l = "<nil>"
	}
	if n.right != nil {
		r = n.right.String()
	} else {
		r = "<nil>"
	}
	return fmt.Sprintf("(%[1]d L: %[2]s - %[1]d R: %[3]s)", n.id, l, r)
}

func main() {
	root := NewNode(0)
	root.right = NewNode(1)
	root.left = NewNode(2)
	root.right.left = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	fmt.Println("original:")
	fmt.Println(root)
	root.Invert()
	fmt.Println("inverted:")
	fmt.Println(root)
}
