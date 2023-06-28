package main

import (
	"fmt"
	"strconv"
)

var fibValues []int64

// fill on the fly
func fillOnFlyFibonacci(n int64) int64 {
	if int64(len(fibValues)) > n {
		return fibValues[n]
	}

	a := fillOnFlyFibonacci(n - 1)
	b := fillOnFlyFibonacci(n - 2)

	fibValues = append(fibValues, a+b)

	return fibValues[n]
}

func preFilledFibonacci(n int64) int64 {
	return fibValues[n]
}

func initSlice() {
	fibValues = make([]int64, 93)
	fibValues[0] = 0
	fibValues[1] = 1

	for i := 2; i < 93; i++ {
		fibValues[i] = fibValues[i-1] + fibValues[i-2]
	}
}

func main() {
	// Fill-on-the-fly.
	// fibValues = make([]int64, 2)
	// fibValues[0] = 0
	// fibValues[1] = 1

	// Prefilled.
	initSlice()

	for {
		// Get n as a string.
		var nString string
		fmt.Printf("N: ")
		fmt.Scanln(&nString)

		// If the n string is blank, break out of the loop.
		if len(nString) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(nString, 10, 64)

		// Uncomment one of the following.
		// fmt.Printf("fibonacci_on_the_fly(%d) = %d\n", n, fillOnFlyFibonacci(n))
		fmt.Printf("fibonacci_prefilled(%d) = %d\n", n, preFilledFibonacci(n))
		// fmt.Printf("fibonacci_bottom_up(%d) = %d\n", n, fibonacci_bottom_up(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibValues); i++ {
		fmt.Printf("%d: %d\n", i, fibValues[i])
	}
}
