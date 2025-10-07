// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер для ограничения времени работы.
package main

import (
	"fmt"
	"sync"
	"time"
)

func writer(timer *time.Timer, wg *sync.WaitGroup, ch chan int) {
	i := 1
	fmt.Println("писатель начал")
	defer wg.Done()
	for {
		select {
		case <-timer.C: // Получили сигнал остановки
			fmt.Println("писатель завершает работу")
			return
		case ch <- i: // Пытаемся записать
			i++
		}
	}

}
func reader(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Println("читатель начал")
	for i := range ch {
		fmt.Printf("%d", i)
	}
}
func main() {
	n := 5
	timer := time.NewTimer(time.Duration(n) * time.Second)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go reader(&wg, ch)
	writer(timer, &wg, ch)

	close(ch)
	wg.Wait()
	fmt.Println("Все горутины завершены, а каналы закрыты.")
}
