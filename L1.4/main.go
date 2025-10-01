package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(signalCh chan os.Signal, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-signalCh:
			fmt.Printf("Worker %d: завершение работы\n", id)
			return
		default:
			i++
			fmt.Printf("Worker %d: выполняет работу, счеткик %d\n", id, i)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(signalCh, i, &wg)
	}

	<-signalCh
	fmt.Println("\nПолучен сигнал завершения...")

	close(signalCh)
	wg.Wait()
	fmt.Println("Все горутины завершены. Выход.")
}
