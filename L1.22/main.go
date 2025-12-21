package main

import (
	"fmt"
	"math/big"
)

func main() {
	var str_num1 string
	var str_num2 string
	x1 := new(big.Int)
	x2 := new(big.Int)
	res := new(big.Int)

	fmt.Println("Мат операции двух очень больших числел")
	fmt.Println("Введите первое число x1 с основанием 10(по условию очень большое)")
	fmt.Scan(&str_num1)
	fmt.Println("Введите втрое число число x2 с основанием 10 (по условию очень большое)")
	fmt.Scan(&str_num2)
	x1.SetString(str_num1, 10)
	x2.SetString(str_num2, 10)
	fmt.Println("вы ввели:\n\t\t", x1, "\n\t\t", x2)
	fmt.Println("x1 * x2 =\t", res.Mul(x1, x2))
	fmt.Println("x1 / x2 =\t", res.Div(x1, x2))
	fmt.Println("x1 + x2 =\t", res.Add(x1, x2))
	fmt.Println("x1 - x2 =\t", res.Sub(x1, x2))
}
