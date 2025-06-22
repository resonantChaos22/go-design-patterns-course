package main

import "fmt"

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewTerminalNode(value int) *Node {
	return &Node{
		Value: value,
	}
}
func NewNode(value int, left, right *Node) *Node {
	n := &Node{
		Value: value,
		left:  left,
		right: right,
	}
	left.parent = n
	right.parent = n
	return n
}

type InorderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func (i *InorderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InorderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}
}

func NewInorderIterator(root *Node) *InorderIterator {
	i := &InorderIterator{
		Current:       root,
		root:          root,
		returnedStart: false,
	}

	//	starting from the leftest node
	for i.Current.left != nil {
		i.Current = i.Current.left
	}

	return i
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) Inorder() *InorderIterator {
	return NewInorderIterator(b.root)
}
func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{
		root: root,
	}
}

func TestTreeTraversal() {
	root := NewNode(1, NewNode(2, NewTerminalNode(4), NewTerminalNode(5)), NewTerminalNode(3))

	it := NewInorderIterator(root)
	for it.MoveNext() {
		fmt.Printf("%d, ", it.Current.Value)
	}
	fmt.Println("\b")

	b := NewBinaryTree(root)
	for i := b.Inorder(); i.MoveNext(); {
		fmt.Printf("%d, ", i.Current.Value)
	}
	fmt.Println("\b")
}
