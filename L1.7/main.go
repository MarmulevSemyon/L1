// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

package main

import (
	"fmt"
	"sync"
)

func Writer_concur_map(m *sync.Map, wg *sync.WaitGroup, num_gorut int) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("горутина %d.%d", num_gorut, i)
		m.Store(key, i)
		fmt.Printf("горутина %d записала %d\n", num_gorut, i)
	}
}
func Writer(m map[string]int, wg *sync.WaitGroup, mutex *sync.Mutex, num_gorut int) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("горутина %d.%d", num_gorut, i)
		m[key] = i
		fmt.Printf("горутина %d записала %d\n", num_gorut, i)
	}
}

func main() {
	m := make(map[string]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	fmt.Println("===с использованием mutex===")

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go Writer(m, &wg, &mutex, i)
	}

	wg.Wait()
	fmt.Println("===с использованием sync.Map===")
	var conc_m sync.Map
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go Writer_concur_map(&conc_m, &wg, i)
	}
	wg.Wait()

}
