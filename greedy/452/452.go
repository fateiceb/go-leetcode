package greedy

import (
	"sort"
)

/*
title: Minimum Number of Arrows to Burst Balloons
There are some spherical balloons spread in two-dimensional space. For each balloon,
provided input is the start and end coordinates of the horizontal diameter.
Since it's horizontal, y-coordinates don't matter,
and hence the x-coordinates of start and end of the diameter suffice.
The start is always smaller than the end.
An arrow can be shot up exactly vertically from different points along the x-axis.
A balloon with xstart and xend bursts by an arrow shot at x if xstart ≤ x ≤ xend.
There is no limit to the number of arrows that can be shot. An arrow once shot keeps traveling up infinitely.
Given an array points where points[i] = [xstart, xend],
return the minimum number of arrows that must be shot to burst all balloons.
Input: points = [[10,16],[2,8],[1,6],[7,12]]
Output: 2
Explanation: One way is to shoot one arrow for example at x = 6 (bursting the balloons [2,8] and [1,6]) and another arrow at x = 11 (bursting the other two balloons).
题目大意：给出了气球直径，对着气球射箭，求最少几支箭能射爆所有气球
思路：按照气球右端点进行升序排序，将箭移动到原本能够射到的最小右端点
这样保证了原本能射到的气球仍然能被射到，比较该右端点于其他气球的左端点
如果右端点比其他气球左端点大，则这只箭也能射到其他几个气球
如果右端点比其他气球左端点小，则说明需要新的箭。
*/

func findMinArrowShots(points [][]int) int {
	//按照右端点排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxright := points[0][1]
	ans := 1
	//遍历points，maxright < p[0]  -> 右端点比其他气球左端点小，则不在这只箭范围内
	//maxright 记录新的右端点 ans++ 箭数+1
	for _, p := range points {
		if maxright < p[0] {
			maxright = p[1]
			ans++
		}
	}
	return ans
}
