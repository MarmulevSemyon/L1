package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Завершение горутины по сигналу из канала")
	var wg sync.WaitGroup
	signal_chan := make(chan int)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-signal_chan:
				wg.Done()
				fmt.Println("горутина закрыта")

				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("\tВ канал был отправлен сигнал об останвки")
	signal_chan <- 1
	wg.Wait()

}
