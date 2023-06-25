package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data       string
	prev, next *Cell
}

func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	after.prev = me

	me.next.prev = after
	me.next = after
}

func (me *Cell) addBefore(before *Cell) {
	me.prev.addAfter(before)
}

func (me *Cell) delete() string {
	next := me.next
	prev := me.prev

	next.prev = prev
	prev.next = next

	return me.data
}

type DoublyLinkedList struct {
	TopSentinel, BottomSentinel *Cell
}

func makeDoublyLinkedList() DoublyLinkedList {
	t := &Cell{data: "Top sentinel"}
	b := &Cell{data: "Bottom sentinel"}

	t.prev = b
	t.next = b

	b.prev = t
	b.next = t

	return DoublyLinkedList{
		TopSentinel:    t,
		BottomSentinel: b,
	}
}

func (dll *DoublyLinkedList) addRange(values []string) {
	lastCell := dll.BottomSentinel

	for _, v := range values {
		next := &Cell{data: v}

		lastCell.addBefore(next)
	}
}

func (dll *DoublyLinkedList) toString(separator string) string {
	cell := dll.TopSentinel
	b := strings.Builder{}

	for cell != dll.BottomSentinel {
		b.WriteString(cell.data)
		b.WriteString(separator)
		cell = cell.next
	}

	return b.String()
}

// fifo
func (dll *DoublyLinkedList) enqueue(data string) {
	dll.TopSentinel.addAfter(&Cell{data: data})
}

func (dll *DoublyLinkedList) pushTop(data string) {
	dll.enqueue(data)
}

func (dll *DoublyLinkedList) dequeue() string {
	return dll.BottomSentinel.prev.delete()
}

func (dll *DoublyLinkedList) popTop() string {
	return dll.dequeue()
}

// lifo
func (dll *DoublyLinkedList) pushBottom(data string) {
	dll.BottomSentinel.addBefore(&Cell{data: data})
}

func (dll *DoublyLinkedList) popBottom() string {
	return dll.BottomSentinel.prev.delete()
}

func (dll *DoublyLinkedList) isEmpty() bool {
	return dll.TopSentinel == dll.BottomSentinel.prev
}

func main() {
	// Make a list from a slice of values.
	list := makeDoublyLinkedList()
	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	list.addRange(animals)
	fmt.Println(list.toString(" "))

	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := makeDoublyLinkedList()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Citrine")
	fmt.Printf("%s ", queue.dequeue())
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.isEmpty() {
		fmt.Printf("%s ", queue.dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := makeDoublyLinkedList()
	deque.pushTop("Ann")
	deque.pushTop("Ben")
	fmt.Printf("%s ", deque.popBottom())
	deque.pushBottom("F-Cat")
	fmt.Printf("%s ", deque.popBottom())
	fmt.Printf("%s ", deque.popBottom())
	deque.pushBottom("F-Dan")
	deque.pushTop("Eva")
	for !deque.isEmpty() {
		fmt.Printf("%s ", deque.popBottom())
	}
	fmt.Printf("\n")
}
