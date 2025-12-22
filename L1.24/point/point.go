package point

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}

}

func (point *Point) Distans(otherPoint Point) float64 {
	dist := math.Sqrt(math.Pow(math.Abs(otherPoint.x-point.x), 2) + math.Pow(math.Abs(otherPoint.y-point.y), 2))
	return dist
}
