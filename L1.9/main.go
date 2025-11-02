package main

import (
	"fmt"
)

func FromArrToChan1(ch1 chan int, arr []int) {
	defer close(ch1)
	for _, v := range arr {
		ch1 <- v
	}
}
func FromChan1ToChan2(ch1 chan int, ch2 chan int) {
	defer close(ch2)
	for v := range ch1 {
		ch2 <- v * v

	}

}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	go FromArrToChan1(ch1, arr)
	go FromChan1ToChan2(ch1, ch2)
	for v := range ch2 {
		fmt.Println(v)
	}
	fmt.Println("канал успешно закрыт")
}
