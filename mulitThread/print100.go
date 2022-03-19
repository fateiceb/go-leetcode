package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 100; i++ {
		ch <- i + 1
	}
}
func consumer(id int, ch chan int, done chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if ok {
			fmt.Printf("id:%d tast:%d\n", id, num)
		} else {
			break
		}
	}
	done <- 1
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	go producer(ch)
	conum := 10
	done := make(chan int, conum)
	wg.Add(conum)
	for i := 0; i < conum; i++ {
		go consumer(i, ch, done, &wg)
	}
	wg.Wait()
	for i := 0; i < conum; i++ {
		<-done
	}
}
