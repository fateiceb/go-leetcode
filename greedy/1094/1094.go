package greedy

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
只有0-1000个车站，记录每个站点，车上的人数，根据容量减去每次车上的人数。
*/
func carPooling(trips [][]int, capacity int) bool {
	//存储当前车站车上的人数
	locations := [1001]int{}
	for i := 0; i < len(trips); i++ {
		p := trips[i]
		locations[p[1]] += p[0]
		locations[p[2]] -= p[0]
	}
	for i := 0; i < len(locations); i++ {
		//容量减去当前车站车上的人
		capacity -= locations[i]
		if capacity < 0 {
			return false
		}
	}

	//wrong
	// sort.Slice(trips, func(i, j int) bool {
	// 	return trips[i][1] < trips[j][1]
	// })
	// right := trips[0][2]
	// rcap := trips[0][0]
	// cap := 0
	// for _, p := range trips {
	// 	//如果该次的右端点一直大于后面的左端点则继续加
	// 	if right > p[1] {
	// 		//cap加上每趟的人数
	// 		cap += p[0]
	// 		//超出上限
	// 		if cap > capacity {
	// 			return false
	// 		}
	// 	} else {
	// 		cap = cap - rcap
	// 		//新的右端点开始
	// 		right = p[2]
	// 		rcap = p[0]
	// 	}
	// }
	return true
}
