package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

// Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой,
// т.е. символы в разных регистрах считать одинаковыми.

// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
func main() {
	fmt.Println(isSet("abcd"), isSet("abCdefA"), isSet("aabcd"), isSet("abcCd"), isSet("qwertyuiop[asdfghjkl;']"))

}
func isSet(str string) bool {
	m := make(map[rune]bool)
	lowerStr := strings.ToLower(str)
	for _, el := range []rune(lowerStr) {
		if m[el] {
			return false
		} else {
			m[el] = true
		}
	}

	return true
}
