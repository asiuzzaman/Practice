package main

import (
	"fmt"
)

func main() {
	s := "helloworld"

	for i := 0; i < 10; i++ {
		fmt.Print(string(s[i]))
	}

}
