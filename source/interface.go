package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{Radius: 6.2}
	rectangle := Rectangle{width: 5.2, height: 10, 5}

	fmt.Println("area circle    : %d\n", getArea(circle))
	fmt.Println("area rectangle : %d\n", getArea(rectangle))

}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Phi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.Area()
}
