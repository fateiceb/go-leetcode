package main

import (
	"fmt"
)

func getSum(a, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func main() {
	fmt.Println(getSum(4, 3))
}
