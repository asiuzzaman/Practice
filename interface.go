package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width, height float64
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func (p Rect) Perimeter() float64 {
	return 2 * (p.width + p.height)
}

func main() {
	s := Rect{5.0, 4.0}
	t := Rect{5.0, 4.0}

	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("value  of t is %v\n", t)
	fmt.Println(t == s)
	fmt.Println(s.Area(), t.Perimeter())

}
