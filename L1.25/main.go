package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}
func main() {
	var s int64
	var wg sync.WaitGroup
	// defer wg.Wait()

	fmt.Println("ВРЕМЯ СПАТЬ\nВведите число секунд на которое хотите усыпить:")
	fmt.Scan(&s)

	go func(wg *sync.WaitGroup, s int64) {
		defer wg.Done()
		for i := range s {
			sleep(time.Second)
			fmt.Printf("прошло %d секунд\n\tвыполняю какую-то работу\n", i+1)
		}
	}(&wg, s)
	sleep(time.Duration(s) * time.Second)
	fmt.Printf("Указанный промежуток времени (%d секунд) прошел", s)

}
