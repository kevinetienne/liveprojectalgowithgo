package main

import (
	"fmt"
	"time"
)

func eulerSieve(max int) []bool {
	sieve := make([]bool, max+1)
	sieve[2] = true

	for i := 3; i <= max; i += 2 {
		sieve[i] = true
	}

	for i := range sieve {
		if sieve[i] {
			maxI := (max / i)
			if maxI%2 == 0 {
				maxI--
			}

			for j := maxI; j >= i; j -= 2 {
				if sieve[i] {
					sieve[i*j] = false
				}
			}
		}
	}

	return sieve
}

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
	sieve1 := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed (sieveOfEratosthenes): %f seconds\n", elapsed.Seconds())

	start = time.Now()
	sieve2 := eulerSieve(max)
	elapsed = time.Since(start)
	fmt.Printf("Elapsed: %f seconds (eulerSieve)\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve1)

		primes := sieveToPrimes(sieve1)
		fmt.Println(primes)

		printSieve(sieve2)

		primes = sieveToPrimes(sieve2)
		fmt.Println(primes)
	}
}
