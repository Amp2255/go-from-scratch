package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Rectangle struct {
	l, b float64
}
type Circle struct {
	radius float64
}

func (rec Rectangle) area() float64 {
	return rec.l * rec.b
}
func (cir Circle) area() float64 {
	return math.Pi * cir.radius * cir.radius
}
func calculate(s Shape) {
	fmt.Println("Area is", s.area())
}
func main() {
	fmt.Println("Welcome to interace!")
	rect := Rectangle{
		l: 2,
		b: 5,
	}
	calculate(rect)
	circle := Circle{
		radius: 6,
	}
	calculate(circle)
}
