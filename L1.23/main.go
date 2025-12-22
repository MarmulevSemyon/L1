package main

import (
	"fmt"
)

func main() {
	fmt.Println("слайсы")
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println("первоначальный слайс:", s)
	fmt.Println("'Удаление элемента с индексом 3'")
	DelElemnt(3, &s)
	fmt.Println("Получившийся слайс:", s)
	fmt.Println("длина слайса", len(s))
	fmt.Println("ёмккость слайса", cap(s))

}

func DelElemnt(index int, s *[]int) {
	if index >= 0 && index < len(*s) {
		*s = append((*s)[:index], (*s)[index+1:]...) // append возвращает новый слайс
		// copy((*s)[index:], (*s)[index+1:]) //
		// *s = (*s)[:len(*s)-1]
	}
}
