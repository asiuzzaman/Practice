package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width, height float64
}

type Circ struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func (c Circ) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Why method is Used?
/*
   Maintain passing parameter is too difficult and Method is used only calling One function
*/
func main() {
	res1 := Rect{5.4, 9.09}
	res2 := Circ{12.5}

	fmt.Println("Ractangle:", res1.Area())
	fmt.Println("Circle:", res2.Area())

}
