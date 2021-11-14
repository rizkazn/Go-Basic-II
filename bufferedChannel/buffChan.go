package main

import (
	"fmt"
	"sync"
)

const min = 0
const max = 100

type mod func(int) bool

func m(x int) bool {
	return x%3 == 0 || x%5 == 0
}

func interval(min, max int) chan int {
	defer func() {
		wg.Done()
	}()
	ch := make(chan int)
	go func() {
		for i := min; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func filter(m mod, ch chan int) chan int {
	ch2 := make(chan int)
	go func() {
		for x := range ch {
			if m(x) {
				ch2 <- x
			}
		}
		close(ch2)
	}()
	return ch2
}

func sum(ch2 chan int) int {
	jumlah := 0
	for i := range ch2 {
		jumlah += i
	}
	return jumlah
}

func result(max int) int {
	return sum(filter(m, interval(min, max)))
}

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)

	fmt.Println("Result from this program : ", result(max))

	wg.Wait()
}
