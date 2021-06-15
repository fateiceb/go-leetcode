/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	maps := make(map[*Node]*Node)
	if head == nil {
		return head
	}
	if val, ok := m[head]; ok {
		return val
	}
	newNode := &Node{
		Val: head.Val,
	}
	m[head] = newNode
	newNode.Next = copyRandomList(head.next)
	newNode.Random = copyRandomList(head.Random)
	return newNode
}