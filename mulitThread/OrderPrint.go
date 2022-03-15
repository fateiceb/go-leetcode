package main

import (
	"fmt"
	"sync"
)

var dog = make(chan struct{}, 1)
var cat = make(chan struct{}, 1)
var done = make(chan struct{}, 1)

var wg sync.WaitGroup

func printDog() {
	wg.Add(1)
	defer wg.Done()
	defer close(dog)
	for i := 0; i < 4; i++ {
		<-cat
		fmt.Println("dog")
		dog <- struct{}{}
	}
}
func printCat() {
	wg.Add(1)
	defer wg.Done()
	defer close(cat)
	for i := 0; i < 4; i++ {
		<-dog
		fmt.Println("cat")
		cat <- struct{}{}
	}
	done <- struct{}{}
}

func main() {
	cat <- struct{}{}
	go printDog()
	go printCat()
	wg.Wait()
	<-done

}
