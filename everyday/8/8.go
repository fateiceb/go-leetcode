package main

import (
	"fmt"
	"strconv"
)

func myAtoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
func main() {
	fmt.Println(myAtoi("12"))
}
