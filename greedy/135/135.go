package greedy

/*
title:candy
content:
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

    Each child must have at least one candy.
    Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.
example:
Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
题目大意：一排孩子，ratings[i]表示第i个孩子的权重，有高权重的孩子拿到的糖果要比邻位的多
解题思路：
*/

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candys[i] = 1
		if i == 0 {
			continue
		}
		if ratings[i] > ratings[i-1] && candys[i] <= candys[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candys[i] <= candys[i+1] {
			candys[i] = candys[i+1] + 1
		}
	}

	sum := 0
	for i := 0; i < len(candys); i++ {
		sum += candys[i]
	}
	return sum
}
