package main

import (
	"fmt"
)

func main() {
	var n1 uint
	var n2 uint

	fmt.Print("Введите размер первого множества(>0)")
	fmt.Scan(&n1)
	set1 := make([]int, n1)

	for i := range int(n1) {
		fmt.Printf("Введите %d-ый элемент первого множества: ", i+1)
		fmt.Scan(&set1[i])
	}

	fmt.Print("Введите размер второго множества (>0)")
	fmt.Scan(&n2)
	set2 := make([]int, n2)

	for i := range int(n2) {
		fmt.Printf("Введите %d-ый элемент второго множества: ", i+1)
		fmt.Scan(&set2[i])
	}
	var res []int
	if n1 > n2 {
		res = make([]int, n2)
	} else {
		res = make([]int, n1)
	}

	hash := make(map[int]bool)
	for _, val := range set1 {
		hash[val] = true
	}
	i := 0
	for _, val := range set2 {
		if hash[val] {
			res[i] = val
			i++
		}
	}

	fmt.Println("пересечение двух последовательностей:")
	fmt.Println(res[0:i])
}
