package main

import (
	"fmt"
)

var s string = "asiuzzaman"

func reverse(n int) {

	if n == 10 {
		return
	}
	n++
	reverse(n)

	fmt.Print(string(s[n-1]))
}

func main() {

	reverse(0)
}
