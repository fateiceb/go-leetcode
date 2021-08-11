package main

import (
	"fmt"
)

func main() {
	array := make([]int, 0, 3)
	array = append(array, 1)
	fmt.Println(array)
}
