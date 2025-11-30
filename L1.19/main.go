package main

import (
	"fmt"
)

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {
	var str string
	fmt.Print("Введите строку, которую хотите перевернуть: ")
	fmt.Scan(&str)

	resStr := reverse(str)
	fmt.Printf("Перевернутая строка:\n\t%s", resStr)

}

func reverse(str string) string {
	resRuneStr := make([]rune, len(str))
	runeStr := []rune(str)

	for i := 0; i < len(runeStr); i++ {
		resRuneStr[i] = runeStr[len(runeStr)-i-1]
	}
	return string(resRuneStr)
}
