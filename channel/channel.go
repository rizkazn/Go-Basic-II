package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go GanjilGenap(ch)
	go fibonacci(ch, 10)

	wg.Wait()
}

func fibonacci(ch chan<- int, n int) {
	defer func() {
		close(ch)
		wg.Done()
	}()

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
}

func GanjilGenap(ch <-chan int) {
	for i := range ch {
		if i%2 != 0 {
			fmt.Println(i, "is ODD")
		} else {
			fmt.Println(i, "is EVEN")
		}
	}
	wg.Done()
}

// func fibonacci(ch chan<- int) {
//     nameString = prosessFibonaci
//     <- sendToGanjilGenap
// }

// func GanjilGenap(ch <-chan int) {
//     result := prosessFilter
//     fmt.Println(“ganjil”, result)
// }
