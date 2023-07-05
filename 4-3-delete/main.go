package main

import (
	"fmt"
	"math/rand"
	"time"
)

type employee struct {
	name, phone string
	deleted     bool
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

		if employee.deleted {
			fmt.Printf("%2d: xxx\n", i)
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
	deletedIndex := -1
	for i := 0; i < lpht.capacity; i++ {
		index := (hash + i) % lpht.capacity

		if lpht.employees[index] == nil {
			if deletedIndex > -1 {
				return deletedIndex, i + 1
			}

			return index, i + 1
		}

		if deletedIndex == -1 && lpht.employees[index].deleted {
			deletedIndex = index
			continue
		}

		if lpht.employees[index].name == name {
			return index, i + 1
		}

	}

	if deletedIndex > -1 {
		return deletedIndex, lpht.capacity
	}

	return -1, lpht.capacity
}

// Add an item to the hash table.
func (lpht *linearProbingHashTable) set(name string, phone string) {
	key, _ := lpht.find(name)
	if key == -1 {
		panic("key not in table")
	}

	if lpht.employees[key] != nil && lpht.employees[key].deleted {
		b := lpht.employees[key]
		b.name = name
		b.phone = phone
		b.deleted = false

		return
	}

	lpht.employees[key] = &employee{name: name, phone: phone}
}

// Return an item from the hash table.
func (lpht *linearProbingHashTable) get(name string) string {
	key, _ := lpht.find(name)
	e := lpht.employees[key]

	if key >= 0 && e != nil && !e.deleted {
		return e.phone
	}

	return ""
}

// Return true if the person is in the hash table.
func (lpht *linearProbingHashTable) contains(name string) bool {
	key, _ := lpht.find(name)

	return key > -1 && lpht.employees[key] != nil && !lpht.employees[key].deleted
}

func (lpht *linearProbingHashTable) delete(name string) {
	key, _ := lpht.find(name)
	if key > -1 && lpht.employees[key] != nil {
		lpht.employees[key].deleted = true
	}
}

// Show this key's probe sequence.
func (lpht *linearProbingHashTable) probe(name string) int {
	// Hash the key.
	hash := hash(name) % lpht.capacity
	fmt.Printf("Probing %s (%d)\n", name, hash)

	// Keep track of a deleted spot if we find one.
	deleted_index := -1

	// Probe up to lpht.capacity times.
	for i := 0; i < lpht.capacity; i++ {
		index := (hash + i) % lpht.capacity

		fmt.Printf("    %d: ", index)
		if lpht.employees[index] == nil {
			fmt.Printf("---\n")
		} else if lpht.employees[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%s\n", lpht.employees[index].name)
		}

		// If this spot is empty, the value isn't in the table.
		if lpht.employees[index] == nil {
			// If we found a deleted spot, return its index.
			if deleted_index >= 0 {
				fmt.Printf("    Returning deleted index %d\n", deleted_index)
				return deleted_index
			}

			// Return this index, which holds nil.
			fmt.Printf("    Returning nil index %d\n", index)
			return index
		}

		// If this spot is deleted, remember where it is.
		if lpht.employees[index].deleted {
			if deleted_index < 0 {
				deleted_index = index
			}
		} else if lpht.employees[index].name == name {
			// If this cell holds the key, return its data.
			fmt.Printf("    Returning found index %d\n", index)
			return index
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deleted_index >= 0 {
		fmt.Printf("    Returning deleted index %d\n", deleted_index)
		return deleted_index
	}

	// There's nowhere to put a new entry.
	fmt.Printf("    Table is full\n")
	return -1
}

// Make a display showing whether each slice entry is nil.
func (lpht *linearProbingHashTable) dumpConcise() {
	// Loop through the slice.
	for i, employee := range lpht.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
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
		{"Ann Archer", "202-555-0101", false},
		{"Bob Baker", "202-555-0102", false},
		{"Cindy Cant", "202-555-0103", false},
		{"Dan Deever", "202-555-0104", false},
		{"Edwina Eager", "202-555-0105", false},
		{"Fred Franklin", "202-555-0106", false},
		{"Gina Gable", "202-555-0107", false},
	}

	hash_table := newLinearProbingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	hash_table.probe("Hank Hardy")
	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hash_table.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	hash_table.dump()

	hash_table.probe("Ann Archer")
	hash_table.probe("Bob Baker")
	hash_table.probe("Cindy Cant")
	hash_table.probe("Dan Deever")
	hash_table.probe("Edwina Eager")
	hash_table.probe("Fred Franklin")
	hash_table.probe("Gina Gable")
	hash_table.set("Hank Hardy", "202-555-0108")
	hash_table.probe("Hank Hardy")

	// Look at clustering.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	big_capacity := 1009
	big_hash_table := newLinearProbingHashTable(big_capacity)
	num_items := int(float32(big_capacity) * 0.9)
	for i := 0; i < num_items; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		big_hash_table.set(str, str)
	}
	big_hash_table.dumpConcise()
	fmt.Printf("Average probe sequence length: %f\n",
		big_hash_table.aveProbeSequenceLength())
}
