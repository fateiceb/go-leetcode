package main

import "fmt"

var one = make(chan struct{}, 1)
var two = make(chan struct{}, 1)
var done = make(chan struct{})

func PrintOne() {
	defer close(one)
	for i := 0; i < 10; i++ {
		//必须由缓冲，不然two被拿走，才可以向下执行
		<-two
		fmt.Println("1")
		one <- struct{}{}
	}
}
func PrintTwo() {
	defer close(two)
	for i := 0; i < 10; i++ {
		<-one
		fmt.Println("2")
		two <- struct{}{}
	}
	done <- struct{}{}
}
func main() {
	defer close(done)
	two <- struct{}{}
	go PrintOne()
	go PrintTwo()
	<-done
}
