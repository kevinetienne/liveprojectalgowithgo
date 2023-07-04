package main

import (
	"fmt"
)

type employee struct {
	name, phone string
}

type chainingHashTable struct {
	numBuckets int
	buckets    [][]*employee
}

func newChainingHashTable(numBuckets int) *chainingHashTable {
	return &chainingHashTable{
		numBuckets: numBuckets,
		buckets:    make([][]*employee, numBuckets),
	}
}

func (ht *chainingHashTable) dump() {
	for i, employees := range ht.buckets {
		fmt.Printf("bucket: %d\n", i)
		for _, employee := range employees {
			fmt.Printf("    %s: %s\n", employee.name, employee.phone)
		}
	}
}

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (ht *chainingHashTable) find(name string) (int, int) {
	n := hash(name)
	key := n % ht.numBuckets
	for i, v := range ht.buckets[key] {
		if v.name == name {
			return key, i
		}
	}

	return key, -1
}

// Add an item to the hash table.
func (ht *chainingHashTable) set(name string, phone string) {
	key, i := ht.find(name)
	if i >= 0 {
		b := ht.buckets[key][i]
		b.name = name
		b.phone = phone

		return
	}
	ht.buckets[key] = append(ht.buckets[key], &employee{name: name, phone: phone})
}

// Return an item from the hash table.
func (ht *chainingHashTable) get(name string) string {
	key, i := ht.find(name)
	if i >= 0 {
		b := ht.buckets[key][i]
		return b.phone
	}

	return ""
}

// Return true if the person is in the hash table.
func (ht *chainingHashTable) contains(name string) bool {
	_, i := ht.find(name)

	return i > -1
}

// Delete this key's entry.
func (ht *chainingHashTable) delete(name string) {
	key, i := ht.find(name)
	if i >= 0 {
		ht.buckets[key] = append(ht.buckets[key][:i], ht.buckets[key][i+1:]...)
	}
}

func hash(s string) int {
	hash := 5381

	for _, c := range s {
		hash = ((hash << 5) + hash) + int(c)
	}

	if hash < 0 {
		hash = -hash
	}

	return hash
}

func main() {
	// Make some names.
	employees := []employee{
		{"Ann Archer", "202-555-0101"},
		{"Bob Baker", "202-555-0102"},
		{"Cindy Cant", "202-555-0103"},
		{"Dan Deever", "202-555-0104"},
		{"Edwina Eager", "202-555-0105"},
		{"Fred Franklin", "202-555-0106"},
		{"Gina Gable", "202-555-0107"},
		{"Herb Henshaw", "202-555-0108"},
		{"Ida Iverson", "202-555-0109"},
		{"Jeb Jacobs", "202-555-0110"},
	}

	ht := newChainingHashTable(10)
	for _, employee := range employees {
		ht.set(employee.name, employee.phone)
	}
	ht.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", ht.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", ht.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	ht.delete("Dan Deever")
	fmt.Printf("Sally Owens: %s\n", ht.get("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", ht.contains("Dan Deever"))
	fmt.Printf("Fred Franklin: %s\n", ht.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	ht.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", ht.get("Fred Franklin"))
}
