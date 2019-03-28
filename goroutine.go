package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//time.Sleep(1 * time.Second)
		go route(i)
	}

	time.Sleep(30 * time.Second)
}

func route(n int) {

	time.Sleep(2 * time.Second)
	fmt.Println(n)

}
