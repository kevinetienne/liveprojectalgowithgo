package main

import "fmt"

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

func main() {
	for {
		var a, b int
		fmt.Print(">>> a: ")
		fmt.Scan(&a)
		fmt.Print(">>> b: ")
		fmt.Scan(&b)
		fmt.Printf("GCD(%d, %d): %d\n", a, b, gcd(a, b))
		fmt.Printf("LCM(%d, %d): %d\n", a, b, lcm(a, b))
		fmt.Println("---")
	}
}
