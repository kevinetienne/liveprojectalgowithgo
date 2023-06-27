package main

import "fmt"

type node struct {
	data        string
	left, right *node
}

func (n *node) insertValue(data string) {
	if n.data > data {
		if n.left == nil {
			n.left = &node{data: data}
			return
		}
		n.left.insertValue(data)
	}

	if n.data < data {
		if n.right == nil {
			n.right = &node{data: data}
			return
		}
		n.right.insertValue(data)
	}
}

func (n *node) findValue(data string) *node {
	if n.data == data {
		return n
	}

	if n.data > data {
		if n.left == nil {
			return nil
		}
		return n.left.findValue(data)
	}

	if n.data < data {
		if n.right == nil {
			return nil
		}
		return n.right.findValue(data)
	}

	return nil
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

func main() {
	// Make a root node to act as sentinel.
	root := node{data: ""}

	// Add some values.
	root.insertValue("I")
	root.insertValue("G")
	root.insertValue("C")
	root.insertValue("E")
	root.insertValue("B")
	root.insertValue("K")
	root.insertValue("S")
	root.insertValue("Q")
	root.insertValue("M")

	// Add F.
	root.insertValue("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.right.inorder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.findValue(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
