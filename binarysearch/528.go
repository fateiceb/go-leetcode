package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// input := 5
	arr := []int{1, 3, 4, 1, 5, 6}
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	rn := rand.Intn(arr[len(arr)-1]) + 1
	k := sort.SearchInts(arr, rn)
	fmt.Printf("%v", arr)
	fmt.Println(rn)
	fmt.Println(k)
}
