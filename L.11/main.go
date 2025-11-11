package main

import (
	"fmt"
)

func main() {
	var n1 uint
	var n2 uint

	fmt.Print("Введите размер первого массива (>0)")
	fmt.Scan(&n1)
	arr1 := make([]int, n1)

	for i := 0; i < int(n1); i++ {
		fmt.Printf("Введите %d-ый элемент первого массива: ", i+1)
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Введите размер второго массива (>0)")
	fmt.Scan(&n2)
	arr2 := make([]int, n2)

	for i := 0; i < int(n2); i++ {
		fmt.Printf("Введите %d-ый элемент второго массива: ", i+1)
		fmt.Scan(&arr2[i])
	}
	var res []int
	if n1 > n2 {
		res = make([]int, n2)
	} else {
		res = make([]int, n1)
	}

	hash := make(map[int]bool)
	for _, val := range arr1 {
		hash[val] = true
	}
	i := 0
	for _, val := range arr2 {
		if hash[val] {
			res[i] = val
			i++
		}
	}

	fmt.Println("пересечение двух последовательностей:")
	fmt.Println(res[0:i])
}
