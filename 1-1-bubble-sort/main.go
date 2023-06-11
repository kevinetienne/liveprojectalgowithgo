package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeRandomArray(numItems, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		arr[i] = r.Intn(max)
	}

	return arr
}

func printArray(arr []int, numItems int) {
	fmt.Println(arr[:numItems])
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

func bubbleSort(arr []int) {
	s := len(arr)
	for i := 0; i < s-1; i++ {
		for j := 0; j < s-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := makeRandomArray(10000, 1000000)
	printArray(arr, 10)
	checkSorted([]int{1, 2, 3, 4, 5})
	checkSorted([]int{5, 43, 3})

	bubbleSort(arr)
	printArray(arr, 10)
	checkSorted(arr)
}
