package main

import (
	"fmt"
)

type My_type struct {
	qwe   int
	qwert string
}

func main() {
	s := "строка"
	i := 123
	b := true
	chan_type := make(chan int)
	some_type := My_type{123, "qwer"}

	fmt.Print(s)
	fmt.Printf(" имеет тип: %s\n", WhatType(s))

	fmt.Print(i)
	fmt.Printf(" имеет тип: %s\n", WhatType(i))

	fmt.Print(b)
	fmt.Printf(" имеет тип: %s\n", WhatType(b))

	fmt.Print(chan_type)
	fmt.Printf(" имеет тип: %s\n", WhatType(chan_type))

	fmt.Print(some_type)
	fmt.Printf(" имеет тип: %s\n", WhatType(some_type))

}

func WhatType(val interface{}) string {
	var s string
	switch val.(type) {
	case int:
		s = "int"
	case string:
		s = "string"
	case bool:
		s = "bool"
	case chan int:
		s = "chan int"
	default:
		s = "Неизвестный тип"
	}

	return s
}
