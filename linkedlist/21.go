pakcage linkedlist
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    h := &ListNode{
    }
    cur := h
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val{
            cur.Next = l1
            l1 = l1.Next
        }else{
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 != nil {
        cur.Next = l1
    }else{
        cur.Next = l2
    }
    return h.Next
}