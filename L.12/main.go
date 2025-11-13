package main

import (
	"fmt"
)

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	// n := 0
	// fmt.Print("Введите размер массива строк(>0)")
	// fmt.Scan(&n)
	// my_arr := make([]int, n)

	// for i := range n {
	// 	fmt.Printf("Введите %d-ый элемент массива: ", i+1)
	// 	fmt.Scan(&my_arr[i])
	// }

	fmt.Print(toSet(arr))

}

func toSet(arr []string) []string {
	set_map := make(map[string]bool)
	for _, val := range arr {
		set_map[val] = true
	}
	set_arr := []string{}
	for key := range set_map {
		set_arr = append(set_arr, key)
	}
	return set_arr
}
