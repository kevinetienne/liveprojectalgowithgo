package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data string
	next *Cell
}

func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	me.next = after
}

func (me *Cell) deleteAfter() {
	if me.next != nil {
		me.next = me.next.next
	}
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	return LinkedList{sentinel: &Cell{data: "start"}}
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

func (ll *LinkedList) length() int {
	count := 0
	cell := ll.sentinel.next

	for cell != nil {
		count += 1
		cell = cell.next
	}

	return count
}

func (ll *LinkedList) isEmpty() bool {
	return ll.sentinel.next == nil
}

func (ll *LinkedList) push(data string) {
	ll.addRange([]string{data})
}

func (ll *LinkedList) pop() string {
	cell := ll.sentinel
	var previous *Cell

	for cell.next != nil {
		previous = cell
		cell = cell.next
	}

	data := previous.next.data
	previous.next = nil

	return data
}

func main() {
	// small_list_test()

	// Make a list from a slice of values.
	greek_letters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := makeLinkedList()
	list.addRange(greek_letters)
	fmt.Println(list.toString(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := makeLinkedList()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.isEmpty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.toString(" "))
	}
}
