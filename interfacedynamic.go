package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func (p Rect) Perimeter() float64 {
	return 2 * (p.width + p.height)
}

func main() {
	var s, t Shape
	s = Rect{5.0, 4.0}
	t = Circle{5.7}
	fmt.Println(t.Area())
	fmt.Println(s.Perimeter())
}
