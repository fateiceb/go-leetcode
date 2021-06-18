type MyCircularQueue struct {
	array []int
	head  int
	tear  int
	n     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		array: make([]int, k+1),
		n:     k + 1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tear = (this.tear + 1) % this.n
	this.array[this.tear] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.n
	return true
}

func (this *MyCircularQueue) Front() int {
	if !this.IsEmpty() {
		return this.array[(this.head+1)%this.n]
	}
	return -1
}

func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.array[this.tear]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tear
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tear+1)%this.n == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */