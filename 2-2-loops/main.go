package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	return LinkedList{sentinel: &Cell{data: "sentinel"}}
}

func (ll *LinkedList) addRange(values []string) {
	lastCell := ll.sentinel
	cell := ll.sentinel

	for cell != nil {
		lastCell = cell
		cell = cell.next
	}

	for _, v := range values {
		next := &Cell{data: v}
		lastCell.next = next
		lastCell = next
	}
}

func (ll *LinkedList) toString(separator string) string {
	cell := ll.sentinel.next
	b := strings.Builder{}

	for cell != nil {
		b.WriteString(cell.data)
		b.WriteString(separator)
		cell = cell.next
	}

	return b.String()
}

func (ll *LinkedList) toStringMax(separator string, max int) string {
	cell := ll.sentinel.next
	b := strings.Builder{}

	for cell != nil && max > 0 {
		b.WriteString(cell.data)
		b.WriteString(separator)
		cell = cell.next
		max -= 1
	}

	return b.String()
}

func (ll *LinkedList) hasLoop() bool {
	fast := ll.sentinel.next
	slow := ll.sentinel.next

	for fast != nil {
		if fast.next == nil {
			return false
		}

		if fast == slow {
			return true
		}

		fast = fast.next.next
		slow = slow.next
	}

	return false
}

func main() {
	// Make a list from a slice of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
