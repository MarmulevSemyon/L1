package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dlina := 10
	arr := make([]int, dlina)
	for i := range arr {
		arr[i] = r.Intn(100)
	}
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println("отсортированный слайс в котором нужно найти элемент:")
	printArr(arr)

	var n int
	fmt.Println("введите число, индекс которого хотите найти:")
	fmt.Scanf("%d", &n)
	ind := binFind(arr, n)
	if ind == -1 {
		fmt.Printf("Элемента <%d> нет в массиве\n", n)
	} else {
		fmt.Printf("Элемент <%d> нахоидится в массиве под индексом %d\n", n, ind)
	}
}
func binFind(arr []int, elem int) int {
	var mid int
	left := 0
	right := len(arr) - 1
	for {
		mid = (left + right) / 2
		fmt.Printf("arr[%d] = %d | ", mid, arr[mid])
		if elem == arr[mid] {
			return mid
		} else if elem < arr[mid] {
			right = mid - 1
		} else if elem > arr[mid] {
			left = mid + 1
		}
		if left > right {
			return -1
		}
	}
}

func printArr(arr []int) {
	for i, el := range arr {
		if i == len(arr)-1 {
			fmt.Printf("[%d]%d\n", i, el)
		} else {
			fmt.Printf("[%d]%d | ", i, el)
		}
	}
}
