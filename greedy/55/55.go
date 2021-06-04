package greedy

/*
title: Jump Game
content:
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.
题目大意：从第0个索引开始跳跃，第[i]个索引代表能跳几步
解题思路：i + nums[i]代表下次跳到的索引位置，
如果i+nums[i] >= len(nums[i]) -1 则到达最终位置
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Constraints:
1 <= nums.length <= 104
0 <= nums[i] <= 105
*/
func canJump(nums []int) bool {
	k := 0 //如果能够到达当前位置，那么也就一定能够到达当前位置的左边所有位置
	for i := 0; i < len(nums); i++ {
		if i > k { //遍历元素位置下标大于当前能够到达的最大位置下标，不能到达
			return false
		}
		//能够到达当前位置，看是否更新能够到达的最大位置k
		k = max(k, i+nums[i])
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
