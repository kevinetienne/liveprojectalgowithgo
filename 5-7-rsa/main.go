package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	min := 10000
	max := 50000

	p := findPrime(min, max, 20)
	q := findPrime(min, max, 20)
	// public key
	n := p * q

	λn := totient(p, q)
	e := randomExponent(λn, 3)
	d := inverseMod(e, λn)

	fmt.Println("*** Public ***")
	fmt.Printf("Public key modulus: %d\n", n)
	fmt.Printf("Public key exponent e: %d\n", e)

	fmt.Println()

	fmt.Println("*** Private ***")
	fmt.Printf("Primes: %d, %d\n", p, q)
	fmt.Printf("λ(n): %d\n", λn)
	fmt.Printf("d: %d\n", d)

	for {
		var m int
		fmt.Printf("\n# Digits: ")
		fmt.Scan(&m)
		if m < 1 {
			break
		}

		// encryption
		ciphertext := fastExpMod(m, e, n)
		fmt.Printf("ciphertext: %d\n", ciphertext)

		// decryption
		plaintext := fastExpMod(ciphertext, d, n)
		fmt.Printf("plaintext: %d\n", plaintext)
	}
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

func isProbablyPrime(num, numOfTests int) bool {
	for i := 0; i < numOfTests; i++ {
		n := randRange(1, num)
		if results := fastExpMod(n, num-1, num); results != 1 {
			return false
		}
	}
	return true
}

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

func inverseMod(a, n int) int {
	t := 0
	newT := 1
	r := n
	newR := a

	for newR != 0 {
		quotient := r / newR
		t, newT = newT, t-quotient*newT
		r, newR = newR, r-quotient*newR
	}

	if r > 1 {
		panic("a is not invertible")
	}
	if t < 0 {
		t = t + n
	}

	return t
}

// Pick a random exponent e in the range (2, λn)
// such that gcd(e, λn) = 1.
func randomExponent(λn, min int) int {
	var e, r int

	for r != 1 {
		e = randRange(min, λn)
		r = gcd(e, λn)
	}

	return e
}

func randRange(min, max int) int {
	d := int64(max - min)
	n, _ := rand.Int(rand.Reader, big.NewInt(d))

	return min + int(n.Int64())
}

func totient(p, q int) int {
	return lcm(p-1, q-1)
}

func gcd(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	if b > a {
		a, b = b, a
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	quotient := a / b
	remainder := a - (b * quotient)

	return gcd(b, remainder)
}

func lcm(a, b int) int {
	g := gcd(a, b)

	return a * (b / g)
}
