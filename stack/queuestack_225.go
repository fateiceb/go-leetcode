package main

import "fmt"

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func queuestackConstructor() MyStack {
	return MyStack{
		[]int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ss := this.queue
	this.queue = ss[:len(ss)-1]

	return ss[len(ss)-1]
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main() {
	mystack := queuestackConstructor()
	mystack.Push(1)
	fmt.Println(mystack)
	mystack.Push(2)
	fmt.Println(mystack)
	mystack.Top()
	fmt.Println(mystack)
	mystack.Pop()
	fmt.Println(mystack)
	fmt.Println(mystack.Empty())
}