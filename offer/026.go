package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	mid := middleNode(head)
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	l1 := head
	meargeList(l1, l2)

}
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func meargeList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1tmp := l1.Next
		l2tmp := l2.Next
		l1.Next = l2
		l1 = l1tmp
		l2.Next = l1
		l2 = l2tmp
	}
}
func createList(arr []int) (list *ListNode) {
	list = &ListNode{}
	cur := list
	for _, v := range arr {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return list.Next
}
func main() {
	arr := []int{1, 2, 3, 4}
	list := createList(arr)
	reorderList(list)

}
