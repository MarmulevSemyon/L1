package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dlina := 10
	arr, arrConst := make([]int, dlina), make([]int, dlina)

	for i := range arr {
		arr[i] = r.Intn(100)
	}

	fmt.Println("\nДикий массив:")
	printArr(arr)

	arr = qsort(arr)

	fmt.Println("\nМассив после вмешательства быстрой сортрировки:")
	printArr(arr)

	fmt.Println("\nМассив из одинаковых элементов до сортировки:")
	printArr(arrConst)

	arrConst = qsort(arrConst)

	fmt.Println("\nМассив из одинаковых элементов после сортировки:")
	printArr(arrConst)

}

func printArr(arr []int) {
	for i, el := range arr {
		if i == len(arr)-1 {
			fmt.Printf("%d\n", el)
		} else {
			fmt.Printf("%d | ", el)
		}
	}
}

func qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivo := arr[0]

	var left, middle, right []int

	for _, elem := range arr {
		if elem == pivo {
			middle = append(middle, elem)
		} else if elem < pivo {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}

	res := append(qsort(left), middle...)
	res = append(res, qsort(right)...)

	return res
}
