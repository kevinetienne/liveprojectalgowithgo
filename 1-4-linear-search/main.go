package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func linearSearch(values []int, target int) (int, int) {
	for i, v := range values {
		if v == target {
			return i, i + 1
		}
	}

	return -1, len(values)
}

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

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	values := randomArr(numItems, max)
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

		i, n := linearSearch(values, t)
		if i == -1 {
			fmt.Printf("Target %s not found, %d tests\n", target, n)
		} else {
			fmt.Printf("values[%d] = %s, %d tests\n", i, target, n)
		}
	}
}
