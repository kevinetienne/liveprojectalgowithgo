package main

import (
	"fmt"
	"math/rand"
	"time"
)

// example taken from grokking
// pick a pivot
// partition the array into two sub arrays
// call quicksort recursively on the two sub arrays
func qsGrokking(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)

	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	less = append(qsGrokking(less), pivot)
	greater = qsGrokking(greater)

	return append(less, greater...)
}

func randomArr(numItems, max int) []int {
	arr := make([]int, numItems)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numItems; i++ {
		arr[i] = r.Intn(max)
	}

	return arr
}

func checkSorted(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}
	fmt.Println("The array is sorted")
}

func quicksort(arr []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	p := partition(arr, lo, hi)
	quicksort(arr, lo, p-1)
	quicksort(arr, p+1, hi)
}

// divides array into two partition
func partition(arr []int, lo, hi int) int {
	pivot := arr[hi] // choose the last element as the pivot

	// temp pivot index
	i := lo - 1

	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			i += 1
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	i += 1
	arr[i], arr[hi] = arr[hi], arr[i]

	return i
}

func main() {
	// quicksort from grokking
	arr := randomArr(100_000_000, 200_000_000)
	fmt.Printf("Array to be sorted: %v\n", arr[:20])

	start := time.Now()
	r := qsGrokking(arr)
	t := time.Since(start)
	fmt.Printf("Time taken: %.2f seconds\n", t.Seconds())

	fmt.Printf("Sorted array: %v\n", r[:20])
	checkSorted(r)

	// workflow
	arr2 := randomArr(100_000_000, 200_000_000)
	fmt.Println("Array to be sorted: ", arr2[:20])

	start = time.Now()
	quicksort(arr2, 0, len(arr2)-1)
	t = time.Since(start)
	fmt.Printf("Time taken: %.2f seconds\n", t.Seconds())

	fmt.Printf("Sorted array: %v\n", arr2[:20])
	checkSorted(arr2)
}
