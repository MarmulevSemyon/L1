package main

import (
	"fmt"

	"L1.24/point"
)

func main() {

	var x1, x2, y1, y2 float64

	fmt.Print("введите координаты первой точки:\n\tx: ")
	fmt.Scan(&x1)
	fmt.Print("\ty: ")
	fmt.Scan(&y1)
	p1 := point.NewPoint(x1, y1)

	fmt.Print("введите координаты второй точки:\n\tx: ")
	fmt.Scan(&x2)
	fmt.Print("\ty: ")
	fmt.Scan(&y2)
	p2 := point.NewPoint(x2, y2)

	fmt.Println("расстояние между точками:", (*p1).Distans(*p2))

}
