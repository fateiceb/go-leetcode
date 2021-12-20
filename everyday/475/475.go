package main

import (
	"fmt"
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) (ans int) {
	sort.Ints(heaters)
	//对于一个位置而言，有距离最近的两个供暖器。
	//在两个供暖器中选择最近的一个，更新ans即可。
	for _, house := range houses {
		//使用二分查询选择最近供暖器
		j := sort.SearchInts(heaters, house+1)
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return ans
}
func main() {
	houses := []int{1, 2, 3, 4}
	heaters := []int{1, 4}
	expect := 1
	fmt.Printf("expect:%d----real:%d---%t", expect,
		findRadius(houses, heaters),
		expect == findRadius(houses, heaters))
}
