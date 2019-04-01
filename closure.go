package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	a := 1
	increment := func() int {
		a++
		return a
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	// closure is here....
	nxeven := evengen()

	fmt.Println(nxeven())
	fmt.Println(nxeven())
	fmt.Println(nxeven())

}

func evengen() func() int {
	i := 50
	return func() int {
		i += 2
		return i
	}
}
