package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Разработать программу, которая переворачивает порядок слов в строке.
// Пример: входная строка: «snow dog sun», выход: «sun dog snow».
// Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать
// дополнительные срезы, а выполнять операцию «на месте».
// 1 2 3 4 5 6

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку, которую хотите перевернуть: ")

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	resStr := reverse(str)
	fmt.Println(reverseWords(resStr))

}

func reverse(str string) []rune {
	resRuneStr := make([]rune, len(str))
	runeStr := []rune(str)

	for i := 0; i < len(runeStr); i++ {
		resRuneStr[i] = runeStr[len(runeStr)-i-1]
	}
	return resRuneStr
}

func reverseWords(strRune []rune) string {
	start := 0
	n := len(strRune)
	for end := 0; end < n; end++ {
		if strRune[end] == ' ' || end == n-1 {
			if end == n-1 {
				end++
			}
			left, right := start, end-1
			for left < right {
				strRune[left], strRune[right] = strRune[right], strRune[left]
				left++
				right--
			}
			start = end + 1
		}
	}

	return string(strRune)
}
