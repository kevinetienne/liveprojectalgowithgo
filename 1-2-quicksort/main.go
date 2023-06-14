package main

import (
	"fmt"
	"math/rand"
	"time"
)

// pick a pivot
// partition the array into two sub arrays
// call quicksoert recursively on the two sub arrays
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

func main() {
	arr := randomArr(200, 2000)
	fmt.Println("Array to be sorted: ", arr[:20])

	r := qsGrokking(arr)
	fmt.Println("Sorted array: ", r[:20])
	checkSorted(r)
}
