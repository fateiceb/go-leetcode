package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func waitconsumer(id int, ch chan int, wg *sync.WaitGroup) {
	for {
		if value, ok := <-ch; ok {
			fmt.Printf("ID:%d,task:%d\n", id, value)
		} else {
			break
		}
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	go producer(ch)
	conum := 2
	wg.Add(conum)
	for i := 0; i < conum; i++ {
		go waitconsumer(i, ch, &wg)
	}
	wg.Wait()
}
