package main

import (
	"fmt"
)

func main() {
	// nums := []int{4, 2, 3, 4}
	// sort.Ints(nums)
	// fmt.Println(nums)
	s := "MCMXCIV"
	// s2 := "MCMXCIV"
	// s3 := "MCMXCVI"
	maps := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	i := 0
	j := 1
	sums := 0
	for i < len(s) {
		if j == len(s) {
			if value, ok := maps[string(s[i])]; ok {
				sums += value
			}
			break
		}
		key := string(s[i]) + string(s[j])
		//key在maps里，则向后两步
		if value, ok := maps[key]; ok {
			sums += value
			i, j = i+2, j+2
		} else {
			if value, ok := maps[string(s[i])]; ok {
				sums += value
				i++
				j++
			}
		}

	}
	fmt.Println(sums)
}
