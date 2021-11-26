// package main

// type LRUCache struct {
// 	Map   map[int]*Node
// 	Cache *List
// 	cap   int
// }

// type Node struct {
// 	Prev *Node
// 	Next *Node
// 	Key  int
// 	Val  int
// }

// type List struct {
// 	Head *Node
// 	Tail *Node
// 	Size int
// }

// func CreateNode(key, val int) *Node {
// 	return &Node{
// 		Key: key,
// 		Val: val,
// 	}
// }
// func InitList() *List {
// 	head := CreateNode(0, 0)
// 	last := CreateNode(0, 0)
// 	head.Next = last
// 	last.Prev = head
// 	return &List{
// 		Head: head,
// 		Tail: last,
// 		Size: 0,
// 	}
// }

// func Constructor(capacity int) LRUCache {
// 	list := InitList()
// 	m := make(map[int]*Node, capacity)
// 	return LRUCache{
// 		Cache: list,
// 		Map:   m,
// 		cap:   capacity,
// 	}
// }
// func (l *List) pushLast(node *Node) {
// 	l.Tail.Next = node
// 	node.Prev = l.Tail
// 	l.Tail = node
// 	l.Size++
// }
// func (l *List) pushFront(node *Node) {
// 	next := l.Head.Next
// 	node.Prev = l.Head
// 	node.Next = next
// 	l.Head.Next = node
// 	next.Prev = node
// 	l.Size++
// }
// func (l *List) removeFirst() int {
// 	key := l.Head.Next.Key
// 	next := l.Head.Next
// 	next.Prev = l.Head
// 	l.Head.Next = next
// 	l.Size--
// 	return key
// }
// func (l *List) remove(node *Node) {
// 	node.Next.Prev = node.Prev
// 	node.Prev.Next = node.Next
// 	l.Size--
// }

// func (this *LRUCache) Get(key int) int {
// 	//如果存在key,更新链表位置，删除node，在最后添加node
// 	//map中不改变
// 	if v, ok := this.Map[key]; ok {
// 		//更新
// 		this.Cache.remove(v)
// 		this.Cache.pushLast(v)
// 		return v.Val
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	node := CreateNode(key, value)
// 	//如果cache的容量和cache的size相同了，则删除最少使用节点，添加新节点到最后
// 	if this.cap == this.Cache.Size {
// 		//删除第一个节点，map中也要删除
// 		rkey := this.Cache.removeFirst()
// 		delete(this.Map, rkey)
// 		this.Cache.pushLast(node)
// 		this.Map[key] = node
// 	} else {
// 		this.Cache.pushLast(node)
// 		this.Map[key] = node
// 	}
// }
// func main() {
// 	lru := Constructor(2)
// 	lru.Put(1, 1)
// 	lru.Put(2, 2)
// 	lru.Get(1)
// 	lru.Put(3, 3)
// 	lru.Get(2)
// }
