package main

import (
	"fmt"
	"math"
)

func fastExp(num, pow int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
		fmt.Println(result, pow, num)
	}

	return result
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

func main() {
	var num, pow, mod int
	for {
		fmt.Print("num: ")
		fmt.Scan(&num)
		if num < 1 {
			break
		}

		fmt.Print("pow: ")
		fmt.Scan(&pow)
		if pow < 1 {
			break
		}

		fmt.Print("mod: ")
		fmt.Scan(&mod)
		if mod < 1 {
			break
		}

		result1 := fastExp(num, pow)
		expected1 := math.Pow(float64(num), float64(pow))
		fmt.Printf("exp: result: %d, expected: %f, %t\n", result1, expected1, float64(result1) == expected1)

		result2 := fastExpMod(num, pow, mod)
		expected2 := int(math.Pow(float64(num), float64(pow))) % mod
		fmt.Printf("expMod: result: %d, expected: %d, %t\n", result2, expected2, result2 == expected2)
	}
}
