package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Customer struct {
	id           string
	numPurchases int
}

func makeRandomSlice(numItems, max int) []Customer {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]Customer, numItems)

	for i := 0; i < numItems; i++ {
		n := r.Intn(max)
		arr[i] = Customer{
			id:           "C" + strconv.Itoa(i),
			numPurchases: n,
		}
	}

	return arr
}

func printSlice(arr []Customer, numItems, limit int) {
	if limit > numItems {
		limit = numItems
	}
	fmt.Println(arr[:limit])
}

func checkSorted(arr []Customer) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].numPurchases > arr[i+1].numPurchases {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}
	fmt.Println("The customer array is sorted")
}

func countingSort(arr []Customer, max int) []Customer {
	c := make([]int, max)

	for _, v := range arr {
		c[v.numPurchases] += 1
	}

	for i := range c {
		if i == 0 {
			continue
		}
		c[i] += c[i-1]
	}

	l := len(arr)
	b := make([]Customer, l)

	for i := range arr {
		v := arr[l-1-i]
		b[c[v.numPurchases]-1] = v
		c[v.numPurchases] -= 1
	}

	return b
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	values := makeRandomSlice(numItems, max)
	printSlice(values, numItems, 40)
	fmt.Println()

	// Sort and display the result.
	start := time.Now()
	sorted := countingSort(values, max)
	t := time.Since(start)
	fmt.Printf("Time taken: %.2f seconds\n", t.Seconds())
	printSlice(sorted, numItems, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
