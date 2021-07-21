package main

import "fmt"

func main() {
	fmt.Println(trailingZeroesTwo(25))
}

func trailingZeroes(n int) int {
	ans := 0
	for i := 5; i <= n; i += 5 {
		currentFactor := i
		for currentFactor%5 == 0 {
			ans++
			currentFactor /= 5
		}
	}
	return ans
}

func trailingZeroesTwo(n int) int {
	ans := 0
	currentMultiple := 5
	for n >= currentMultiple {
		ans += (n / currentMultiple)
		currentMultiple *= 5
	}
	return ans
}
