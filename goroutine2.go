package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var cnt int = 0

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayhello()
		go increment()

	}
	wg.Wait()

}

func sayhello() {
	fmt.Println("Hello", cnt)
	wg.Done()
}

func increment() {

	cnt++
	wg.Done()

}
