package greedy

/*
title:Gas Station
content:
There are n gas stations along a circular route,
where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and
it costs cost[i] of gas to travel from the ith
station to its next (i + 1)th station. You begin
the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the
starting gas station's index if you can travel
around the circuit once in the clockwise direction,
otherwise return -1. If there exists a solution,
it is guaranteed to be unique
Example:
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit
of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough
to travel back to station 3.
Therefore, return 3 as the starting index.


Constraints:
    gas.length == n
    cost.length == n
    1 <= n <= 104
    0 <= gas[i], cost[i] <= 104
题目大意：环形公路，汽车有一个无限的油箱，gas[i]代表能加多少油,
cost i代表第i个地方的损耗
解题思路：将cost与gas索引统一,先选择能跑路的起始点，

*/
func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; i++ {
		j := i
		remain := gas[i]
		for (remain - cost[j]) >= 0 {
			remain = remain - cost[j] + gas[(j+1)%l]
			j = (j + 1) % l
			if j == i {
				return i
			}
		}
	}

	return -1
}
