package main

import (
	"fmt"
	"sync"
)

type deretBilang struct {
	N int
}

// deret bilangan prima
func (d deretBilang) prima() {
	primes := make([]int, 0, d.N)
	for i := 2; i < d.N; i++ {
		isPrime := true
		for _, j := range primes {
			if i%j == 0 {
				isPrime = false
				continue
			}
		}
		if isPrime == true {
			primes = append(primes, i)

		}
	}
	fmt.Println("Deret Bilangan Prima: ", primes)
	wg.Done()
}

// deret bilangan ganjil
func (d deretBilang) ganjil() {
	odds := make([]int, 0, d.N)
	for i := 0; i < d.N; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		} else {
			continue
		}
	}
	fmt.Println("Deret Bilangan Ganjil: ", odds)
	wg.Done()
}

// deret bilangan genap
func (d deretBilang) genap() {
	evens := make([]int, 0, d.N)
	for i := 0; i < d.N; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		} else {
			continue
		}
	}
	fmt.Println("Deret Bilangan Genap: ", evens)
	wg.Done()
}

// deret bilangan fibonacci
func (d deretBilang) fibonacci() {
	fibos := make([]int, d.N+1, d.N+2)
	if d.N < 2 {
		fibos = fibos[0:2]
	}
	fibos[0] = 0
	fibos[1] = 1
	for i := 2; i <= d.N; i++ {
		fibos[i] = fibos[i-2] + fibos[i-1]
	}
	fmt.Println("Deret Bilangan Fibonacci: ", fibos)
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	var deret = deretBilang{40}
	wg.Add(4)

	fmt.Println("2. Buatlah sebuah struct berisi method yang mengolah object struct N sebagai limit untuk mencetak deret bilangan prima, ganjil/genap, dan fibonacci")
	go deret.prima()
	go deret.ganjil()
	go deret.genap()
	go deret.fibonacci()

	wg.Wait()
}
