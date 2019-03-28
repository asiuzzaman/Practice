package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // synchronizing different go routine. That's why we use Go routine.

func main() {
	var msg = "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)

	}(msg)
	msg = "GoodBye"
	fmt.Println(msg)
	wg.Wait()
}
