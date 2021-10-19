package main

import "fmt"

func findComplement(num int) int {
	highBit := 0
	for i := 1; i <= 30; i++ {
		if 1<<i > num {
			break
		}
		highBit = i
	}
	fmt.Println(highBit)
	mask := 1<<(highBit+1) - 1
	return num ^ mask
}

func main() {
	fmt.Println(1<<3 - 1)
	findComplement(5)
}
