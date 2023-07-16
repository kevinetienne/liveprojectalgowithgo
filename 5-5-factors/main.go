package main

import (
	"fmt"
	"time"
)

func findFactors(num int) []int {
	results := make([]int, 0)

	for num%2 == 0 {
		results = append(results, 2)
		num /= 2
	}

	factor := 3
	for factor*factor <= num {
		if num%factor == 0 {
			results = append(results, factor)
			num /= factor
		} else {
			factor += 2
		}
	}

	if num > 1 {
		results = append(results, num)
	}

	return results
}

func findFactorsSieve(num int, primes []int) []int {
	results := make([]int, 0)

	for num%2 == 0 {
		results = append(results, 2)
		num /= 2
	}

	for _, prime := range primes[1:] {
		if num%prime == 0 {
			results = append(results, prime)
			num /= prime
		} else if prime*prime > num {
			break
		}
	}

	if num > 1 {
		results = append(results, num)
	}

	return results
}

func multiplySlice(factors []int) int {
	t := 1
	for _, x := range factors {
		t *= x
	}

	return t
}

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
	var num int

	s := eulerSieve(20_000_000)
	primes := sieveToPrimes(s)

	for {
		fmt.Print("factor: ")
		fmt.Scan(&num)

		if num < 2 {
			break
		}

		// Find the factors the slow way.
		start := time.Now()
		factors := findFactors(num)
		elapsed := time.Since(start)
		fmt.Printf("find_factors:       %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = findFactorsSieve(num, primes)
		elapsed = time.Since(start)
		fmt.Printf("find_factors_sieve: %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()

	}
}
