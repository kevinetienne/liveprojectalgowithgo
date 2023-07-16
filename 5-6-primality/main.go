package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

const numTests = 20

func fastExpMod(num, pow, mod int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
			result %= mod
		}
		pow /= 2
		num *= num
		num %= mod
	}

	return result
}

func randRange(min, max int) int {
	d := int64(max - min)
	n, _ := rand.Int(rand.Reader, big.NewInt(d))

	return min + int(n.Int64())
}

func isProbablyPrime(num, numOfTests int) bool {
	for i := 0; i < numOfTests; i++ {
		n := randRange(1, num)
		if results := fastExpMod(n, num-1, num); results != 1 {
			return false
		}
	}
	return true
}

func findPrime(min, max, numTests int) int {
	for {
		p := randRange(min, max)
		if p%2 == 0 {
			continue
		}

		if isProbablyPrime(p, numTests) {
			return p
		}
	}
}

func testKnownValues() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Printf("Probability: %f%%\n\n", (1.0-1.0/math.Pow(2, numTests))*100.0)

	fmt.Println("Primes:")
	for _, p := range primes {
		if isProbablyPrime(p, numTests) {
			fmt.Printf("%d Prime\n", p)
		} else {
			fmt.Printf("%d Composite\n", p)
		}
	}

	fmt.Println("Composites:")
	for _, c := range composites {
		if isProbablyPrime(c, numTests) {
			fmt.Printf("%d Prime\n", c)
		} else {
			fmt.Printf("%d Composite\n", c)
		}
	}

}

func main() {
	testKnownValues()

	// Generate random primes.
	for {
		// Get the number of digits.
		var numDigits int
		fmt.Printf("\n# Digits: ")
		fmt.Scan(&numDigits)
		if numDigits < 1 {
			break
		}

		// Calculate minimum and maximum values.
		min := int(math.Pow(10.0, float64(numDigits-1)))
		max := 10 * min
		if min == 1 {
			min = 2
		} // 1 is not prime.

		// Find a prime.
		fmt.Printf("Prime: %d\n", findPrime(min, max, numTests))
	}
}
