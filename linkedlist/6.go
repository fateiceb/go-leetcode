package linkedlist

/*
title:剑指 Offer 06. 从尾到头打印链表
content:
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) (arr []int) {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		arr = append(arr, stack[i])
	}
	return arr
}
