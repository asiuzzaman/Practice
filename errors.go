package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cannot work with 42")
	}

	return arg + 3, nil
}

func main() {

	A := []int{7, 55, 89, 1, 42}

	for _, i := range A {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
}
