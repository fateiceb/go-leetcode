package linkedlist

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList() *ListNode {
	return &ListNode{
		Val:  1,
		Next: nil,
	}
}

func (l *ListNode) addNode(val int) {
	tear := l
	for tear.Next != nil {
		tear = tear.Next
	}
	tear.Next = &ListNode{
		Val:  val,
		Next: nil,
	}

}
