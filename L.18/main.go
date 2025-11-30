package main

import (
	"fmt"
	"sync"
)

type Count struct {
	count int
	sync.Mutex
}

func (c *Count) Increment() {
	defer c.Unlock()
	c.Lock()
	c.count++
}

func (c *Count) GetVal() int {
	return c.count
}

func main() {
	c := Count{count: 0}
	wg := sync.WaitGroup{}

	fmt.Printf("В начале main\tСчетчик равен %d\n\n", c.count)

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int, counter *Count) {
			defer wg.Done()
			fmt.Printf("Горутина №%d инкрементирует счетчик\n", i)
			counter.Increment()
			fmt.Printf("\tСчетчик равен %d\n", counter.GetVal())

		}(&wg, i, &c)
	}
	wg.Wait()

	fmt.Printf("\nВ конце main\tСчетчик равен %d\n", c.GetVal())
}
