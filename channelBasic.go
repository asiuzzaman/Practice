package main

import "fmt"

func square(p chan int) {
	for i := 0; i < 9; i++ {
		p <- i * i
	}
	close(p) // closing channel is a must
}

func main() {

	fmt.Println("Main() func started")
	c := make(chan int, 1)
	go square(c)
	for val := range c {
		fmt.Println(val)
	}
	var s int
	fmt.Scan(&s)
	fmt.Println("Main() Stopped")

}
