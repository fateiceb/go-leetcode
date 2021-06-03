package greedy

import "sort"

/*
title: Car Pooling
content:
You are driving a vehicle that has capacity empty seats initially
available for passengers.  The vehicle only drives east
 (ie. it cannot turn around and drive west.)

Given a list of trips, trip[i] = [num_passengers, start_location,
end_location] contains information about the i-th trip:
the number of passengers that must be picked up,
and the locations to pick them up and drop them off.
The locations are given as the number of kilometers due east from
your vehicle's initial location.

Return true if and only if it is possible to pick up and drop off
all passengers for all the given trips.
Input: trips = [[2,1,5],[3,3,7]], capacity = 4
Output: false
题目大意：trips[i] = [乘客数量，起始位置，结束位置]，从正东方向接乘客放乘客
如果每趟trips都可以完成，则返回true否则返回false
解题思路：
记录每一趟的右端点，如果右端点大于下一趟的左端点，则capicity算在一起
否则记录下一趟的右端点，继续比较下下趟的左端点。
*/
func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][2] < trips[j][2]
	})
	right := trips[0][2]
	cap := 0
	for _, p := range trips {
		//如果该次的右端点一直大于后面的左端点则继续加
		if right > p[1] {
			//cap加上每趟的人数
			cap += p[0]
			//超出上限
			if cap > capacity {
				return false
			}
		} else {
			//新的右端点开始
			right = p[2]
		}
	}
	return true
}
