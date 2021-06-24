package main

import "fmt"

/*
剑指 Offer 15. 二进制中1的个数
请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
输入必须是长度为 32 的 二进制串


*/
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		last := num % 2
		if last == 1 {
			cnt++
		}
		num = num / 2
	}
	return cnt
}
func hammingWeight2(num uint32) int {
	cnt := 0
	for ; num > 0; num &= num - 1 {
		cnt++
	}
	return cnt
}
func main() {
	var num uint32
	num = 0b11111111111111111111111111111001
	fmt.Println(hammingWeight2(num))
}
