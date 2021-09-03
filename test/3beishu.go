package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	m := make(map[int]bool, 0)
	i := 0
	cnt := 0
	for len(arr)-1 != len(m) {
		if !m[i] {
			cnt++
		}
		if cnt%3 == 0 {
			m[i] = true
		}
		if i == len(arr)-1 {
			i = 0
		} else {
			i++
		}
	}

	for i := range arr {
		if m[i] != true {
			fmt.Println(arr[i])
		}
	}
}

// 0 1 2 3 4
// 0 1 2 3
