package main

import (
	"fmt"
	"math/rand"
	"time"
)

type employee struct {
	name, phone string
}

type linearProbingHashTable struct {
	capacity  int
	employees []*employee
}

// Initialize a linearProbingHashTable and return a pointer to it.
func newLinearProbingHashTable(capacity int) *linearProbingHashTable {
	return &linearProbingHashTable{
		capacity:  capacity,
		employees: make([]*employee, capacity),
	}
}

func (lpht *linearProbingHashTable) dump() {
	for i, employee := range lpht.employees {
		if employee == nil {
			fmt.Printf("%2d: ---\n", i)
			continue
		}
		fmt.Printf("%2d: %15s %s\n", i, employee.name, employee.phone)
	}
}

// Return the key's index or where it would be if present and
// the probe sequence length.
// If the key is not present and the table is full, return -1 for the index.
func (lpht *linearProbingHashTable) find(name string) (int, int) {
	hash := hash(name) % lpht.capacity
	for i := 0; i < lpht.capacity; i++ {
		index := (hash + i) % lpht.capacity

		if lpht.employees[index] == nil {
			return index, i + 1
		}

		if lpht.employees[index].name == name {
			return index, i + 1
		}

	}

	return -1, lpht.capacity
}

// Add an item to the hash table.
func (lpht *linearProbingHashTable) set(name string, phone string) {
	key, _ := lpht.find(name)
	if key == -1 {
		panic("key not in table")
	}

	if lpht.employees[key] != nil {
		b := lpht.employees[key]
		b.name = name
		b.phone = phone

		return
	}

	lpht.employees[key] = &employee{name: name, phone: phone}
}

// Return an item from the hash table.
func (lpht *linearProbingHashTable) get(name string) string {
	key, _ := lpht.find(name)
	e := lpht.employees[key]

	if key >= 0 && e != nil {
		return e.phone
	}

	return ""
}

// Return true if the person is in the hash table.
func (lpht *linearProbingHashTable) contains(name string) bool {
	key, _ := lpht.find(name)

	return key > -1 && lpht.employees[key] != nil
}

// Make a display showing whether each slice entry is nil.
func (lpht *linearProbingHashTable) dumpConcise() {
	// Loop through the slice.
	for i, employee := range lpht.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (lpht *linearProbingHashTable) aveProbeSequenceLength() float32 {
	totalLength := 0
	numValues := 0
	for _, employee := range lpht.employees {
		if employee != nil {
			_, probeLength := lpht.find(employee.name)
			totalLength += probeLength
			numValues++
		}
	}
	return float32(totalLength) / float32(numValues)
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
	}

	hashTable := newLinearProbingHashTable(10)
	for _, employee := range employees {
		hashTable.set(employee.name, employee.phone)
	}
	hashTable.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hashTable.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	// fmt.Println("Deleting Dan Deever")
	// hash_table.delete("Dan Deever")
	// fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hashTable.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hashTable.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))

	// Look at clustering.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	bigCapacity := 1009
	bigHashTable := newLinearProbingHashTable(bigCapacity)
	numItems := int(float32(bigCapacity) * 0.9)
	for i := 0; i < numItems; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		bigHashTable.set(str, str)
	}
	bigHashTable.dumpConcise()
	fmt.Printf("Average probe sequence length: %f\n",
		bigHashTable.aveProbeSequenceLength())
}
