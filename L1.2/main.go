package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	s := []int{2, 4, 6, 8, 10}
	wg.Add(5)
	for _, value := range s {
		go func(v int) {
			defer wg.Done()
			fmt.Printf("%d\n", v*v)
		}(value)
	}

	wg.Wait()
}
