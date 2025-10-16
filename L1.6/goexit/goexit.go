package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func workerWithGoexit(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {

		if i == 7 {
			fmt.Printf("Вызывается runtime.Goexit() на итерации %d\n", i)
			runtime.Goexit()
		}

		fmt.Printf("горутина работает:\n \tитерация %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("\tЗавершение горутины по условию с помощью runtime.Goexit()")

	wg.Add(1)
	go workerWithGoexit(&wg)

	wg.Wait()
	defer fmt.Println("горутина завершена")

}
