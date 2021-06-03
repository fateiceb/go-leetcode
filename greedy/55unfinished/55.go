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
	numsIndexLen := len(nums) - 1
	i := 0
	for i < numsIndexLen {
		if nums[i] == 0 {
			return false
		}
		for j := nums[i]; j > 0; j-- {
			if i+j < numsIndexLen && nums[i+j] != 0 {
				i = i + j
				break
			}
			if j == 1 {
				return false
			}
		}

	}
	return true
}
