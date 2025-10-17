package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a int64
	var i uint16
	var bit uint

	for {
		fmt.Print("Введите число типа int64, которому будете менять i-ый бит: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		val, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Printf("err: %s\nпопробуй ещё\n", err)
		} else {
			a = val
			break
		}
	}

	for {
		fmt.Print("номер бита, которого хотите поменять (от 1 до 64 включительно): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		val, err := strconv.ParseUint(input, 10, 16)
		if err != nil {
			fmt.Printf("err: %s\nпопробуй ещё\n", err)

		} else if val > 64 || val == 0 {
			fmt.Printf("i слишком большой или равен нулю О_о\nпопробуй ещё\n")

		} else {
			i = uint16(val) - 1
			break
		}
	}

	for {
		fmt.Print("Введите значение бита (0 или 1): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		val, err := strconv.ParseUint(input, 10, 8)
		if err != nil {
			fmt.Printf("err: %s\nпопробуй ещё\n", err)
		} else if val != 0 && val != 1 {
			fmt.Println("Бит должен быть равен 0 или 1\nПопробуй ещё")
		} else {
			bit = uint(val)
			break
		}
	}

	fmt.Printf("До:\t%d = %064b\n", a, uint64(a))

	if bit == 1 {
		a |= 1 << i
	} else {
		a &^= 1 << i
	}

	fmt.Printf("После:\t%d = %064b\n", a, uint64(a))
}
