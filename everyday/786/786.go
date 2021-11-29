package main

import (
	"fmt"
	"sort"
)

type Element struct {
	Molecule       int
	MoleculeIdx    int
	Denominator    int
	DenominatorIdx int
	result         float64
}

func (e *Element) Cal() bool {
	if e.Denominator != 0 {
		e.result = float64(e.Molecule) / float64(e.Denominator)
		return true
	}
	return false
}
func kthSmallestPrimeFraction(arr []int, k int) []int {
	results := make([]Element, 0)
	for i := 0; i < len(arr); i++ {
		molecule := arr[i]
		for j := len(arr) - 1; j > i; j-- {
			denominator := arr[j]
			ele := &Element{molecule, i, denominator, j, 0}
			if !ele.Cal() {
				return []int{}
			}
			results = append(results, *ele)
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].result < results[j].result
	})
	fmt.Printf("%+v", results)
	return []int{results[k-1].MoleculeIdx, results[k-1].DenominatorIdx}
}
func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	kthSmallestPrimeFraction(arr, k)
}
