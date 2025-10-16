package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Завершение горутины по условию")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина работает, пока счетчик не будет равен 5 (по условию)")

		for i := 1; i <= 5; i++ {
			fmt.Printf("\t счетчик = %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("Горутина закончила работу")

}
