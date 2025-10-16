package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Воркер %d начал работу\n", id)

	for {
		select {

		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу: %v\n", id, ctx.Err())
			return

		default:
			fmt.Printf("Воркер %d выполняет работу...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("\tЗавершение горутины по условию с помощью контекста")

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx, &wg, 1)

	time.Sleep(3 * time.Second)

	fmt.Println("\tОтправляем сигнал на завершение всем воркерам")
	cancel() // Отменяем контекст

	wg.Wait()
	fmt.Println("Все горутины успешно завершены")
}
