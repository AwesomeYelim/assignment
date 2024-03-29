package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go Square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
		fmt.Println("Square")
	}
	//close(ch)

	wg.Wait()
}

func Square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Square:", n*n)
		time.Sleep(time.Second)
	}
	fmt.Println("Done")
	wg.Done()
}
