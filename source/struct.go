package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	var c Circle
	c = Circle{1, 3, 5}
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5

	fmt.Println(circleArea(c))
}
