package main

import (
	"fmt"
)

type node struct {
	data        string
	left, right *node
}

func buildTree() *node {
	a := node{data: "A"}
	b := node{data: "B"}
	c := node{data: "C"}
	d := node{data: "D"}
	e := node{data: "E"}
	f := node{data: "F"}
	g := node{data: "G"}
	h := node{data: "H"}
	i := node{data: "I"}
	j := node{data: "J"}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f
	e.left = &g
	f.left = &h
	h.left = &i
	h.right = &j

	return &a
}

func (n *node) displayIndented(depth int) string {
	format := fmt.Sprintf("%%%ds\n", depth)
	result := fmt.Sprintf(format, n.data)

	if n.left != nil {
		result += n.left.displayIndented(depth + 2)
	}

	if n.right != nil {
		result += n.right.displayIndented(depth + 2)
	}

	return result
}

func (n *node) preorder() string {
	result := n.data

	if n.left != nil {
		result += " "
		result += n.left.preorder()
	}

	if n.right != nil {
		result += " "
		result += n.right.preorder()
	}

	return result
}

func (n *node) inorder() string {
	result := ""

	if n.left != nil {
		result += n.left.inorder()
		result += " "
	}

	result += n.data

	if n.right != nil {
		result += " "
		result += n.right.inorder()
	}

	return result
}

func (n *node) postorder() string {
	result := ""

	if n.left != nil {
		result += n.left.postorder()
		result += " "
	}

	if n.right != nil {
		result += n.right.postorder()
		result += " "
	}

	result += n.data

	return result
}

func (n *node) breadthFirst() string {
	result := ""
	queue := makeDll[*node]()
	queue.enqueue(n)

	for !queue.isEmpty() {
		nn := queue.dequeue()
		result += nn.data

		if nn.left != nil {
			queue.enqueue(nn.left)
		}
		if nn.right != nil {
			queue.enqueue(nn.right)
		}

		if !queue.isEmpty() {
			result += " "
		}
	}

	return result
}

type doubleLinkedList[T any] struct {
	top, bottom *cell[T]
}

type cell[T any] struct {
	prev, next *cell[T]
	data       T
}

func makeDll[T any]() doubleLinkedList[T] {
	top := &cell[T]{}
	bottom := &cell[T]{}

	top.next = bottom
	top.prev = bottom

	bottom.next = top
	bottom.prev = top

	return doubleLinkedList[T]{
		top:    top,
		bottom: bottom,
	}
}

func (dll *doubleLinkedList[T]) enqueue(data T) {
	cell := &cell[T]{
		data: data,
		prev: dll.top,
		next: dll.top.next,
	}
	dll.top.next.prev = cell
	dll.top.next = cell
}

func (dll *doubleLinkedList[T]) dequeue() T {
	toDelete := dll.bottom.prev
	toDelete.prev.next = toDelete.next
	toDelete.next.prev = toDelete.prev

	return toDelete.data
}

func (dll *doubleLinkedList[T]) toString(separator string) string {
	cell := dll.top.next

	result := ""
	for cell != dll.bottom {
		result += fmt.Sprint(cell.data)

		cell = cell.next

		if cell != dll.bottom {
			result += fmt.Sprint(separator)
		}
	}

	return result
}

func (dll *doubleLinkedList[T]) isEmpty() bool {
	return dll.top == dll.bottom.prev
}

func main() {
	root := buildTree()
	fmt.Print(root.displayIndented(0))
	fmt.Println("Preorder:         ", root.preorder())
	fmt.Println("Inorder:          ", root.inorder())
	fmt.Println("Postorder:        ", root.postorder())
	fmt.Println("BreadthFirst:     ", root.breadthFirst())
}
