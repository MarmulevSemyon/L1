package main

import (
	"flag"
	"fmt"
	"sync"
	// "chan"
)

// Работа нескольких воркеров
// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	ch := make(chan int)
	defer close(ch)

	n := flag.Int("n", 0, "количество воркеров")

	flag.Parse()

	wg.Add(*n)
	for range *n {
		go worker(&wg, ch)
	}

	for i := range 100 {
		ch <- i
	}
}
func worker(wg *sync.WaitGroup, ch chan int) {

	for num := range ch {
		fmt.Println(num)
	}
	wg.Done()
}
