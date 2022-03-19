package main

import "fmt"

func Producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id %d,value %d \n", id, value)
		} else {
			break
		}
	}
	done <- true
}
func main() {
	ch := make(chan int, 3)
	conum := 2
	done := make(chan bool, conum)
	for i := 0; i < conum; i++ {
		go Consumer(i, ch, done)
	}
	go Producer(ch)
	for i := 0; i < conum; i++ {
		<-done
	}
}
