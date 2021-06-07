package greedy

import (
	"sort"
)

/*
title:Queue Reconstruction by Height
content:
You are given an array of people, people, which are the attributes
of some people in a queue (not necessarily in order).
Each people[i] = [hi, ki] represents the ith person of height
hi with exactly ki other people in front who have a height
greater than or equal to hi.

Reconstruct and return the queue that is represented by the
input array people. The returned queue should be formatted as an
array queue, where queue[j] = [hj, kj] is the attributes of the jth
person in the queue (queue[0] is the person at the front of the queue).

*/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i int, j int) bool {
		return (people[i][0] > people[j][0]) ||
			(people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	ans := make([][]int, 2000)
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)

	}
	return people
}
