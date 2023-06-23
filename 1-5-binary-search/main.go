package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randomArr(numItems, max int) []int {
	arr := make([]int, numItems)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < numItems; i++ {
		arr[i] = r.Intn(max)
	}

	return arr
}

func printArray(arr []int, numItems, limit int) {
	if limit > numItems {
		limit = numItems
	}
	fmt.Println(arr[:limit])
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

func binarySearch(values []int, target int) (int, int) {
	numTests := 0
	min := 0
	max := len(values) - 1

	for {
		numTests += 1

		if max < min {
			return -1, len(values)
		}

		i := max - ((max - min) / 2)
		guess := values[i]
		if guess == target {
			return i, numTests
		}

		if guess < target {
			min = i + 1
		} else {
			max = i - 1
		}
	}
}

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	values := randomArr(numItems, max)
	printArray(values, numItems, 40)
	quicksort(values, 0, numItems-1)
	printArray(values, numItems, 40)

	for {
		var target string
		fmt.Printf("Target: ")
		fmt.Scanln(&target)
		if target == "" {
			break
		}

		t, err := strconv.Atoi(target)
		if err != nil {
			fmt.Printf("%s is not a valid number\n", target)
			continue
		}

		i, n := binarySearch(values, t)
		if i == -1 {
			fmt.Printf("Target %s not found, %d tests\n", target, n)
		} else {
			fmt.Printf("values[%d] = %s, %d tests\n", i, target, n)
		}
	}
}
