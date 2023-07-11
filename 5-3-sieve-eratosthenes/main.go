package main

import (
	"fmt"
	"time"
)

func sieveOfEratosthenes(max int) []bool {
	sieve := make([]bool, max+1)
	for i := 2; i < len(sieve); i++ {
		sieve[i] = true
	}

	for i := range sieve {
		if sieve[i] {
			for j := i + i; j < len(sieve); j += i {
				sieve[j] = false
			}
		}
	}

	return sieve
}

func printSieve(sieve []bool) {
	i := 2
	for i < len(sieve) {
		if sieve[i] {
			fmt.Printf("%d ", i)
		}
		if i > 2 {
			i++
		}
		i++
	}
	fmt.Println()
}

func sieveToPrimes(sieve []bool) []int {
	var primes []int
	i := 2
	for i < len(sieve) {
		if sieve[i] {
			primes = append(primes, i)
		}
		if i > 2 {
			i++
		}
		i++
	}

	return primes
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve)

		primes := sieveToPrimes(sieve)
		fmt.Println(primes)
	}
}
