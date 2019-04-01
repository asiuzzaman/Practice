package main

import "fmt"

// Multiple Embadding is here

type NameObj struct {
	Name string
}
type Shape struct {
	NameObj   // Inheritance
	color     int32
	isRegular bool
}

type Point struct {
	x, y float64
}

type Rectangular struct {
	NameObj /// Multiple Inheritance is here
	Shape
	center        Point
	width, height float64
}

func main() {

	obj := Rectangular{
		NameObj{"Asiuzzaman"},
		Shape{NameObj{"DefineShape"}, 66, true},
		Point{5, 7},
		100, 200,
	}

	fmt.Println(obj.Shape.NameObj.Name)
	fmt.Println(obj.Shape.isRegular)

}
