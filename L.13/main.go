package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6
	fmt.Printf("\nДо обмена: x=%d,y=%d\n", x, y)
	// fmt.Printf("\nx = x^y:\n%s ^ %s\n", strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2))
	x = x ^ y
	// fmt.Printf("\ny = x^y:\n0%s ^ %s\n", strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2))
	y = x ^ y
	// fmt.Printf("\nx = x^y:\n0%s ^ %s\n", strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2))
	x = x ^ y
	fmt.Printf("\nПосле обмена: x=%d,y=%d\n", x, y)

	/////////////////////////////////
	fmt.Printf("\nДо обмена: x=%d,y=%d\n", x, y)
	x = x + y
	y = x - y
	x = x - y
	fmt.Printf("После обмена: x=%d, y=%d\n", x, y)
}
