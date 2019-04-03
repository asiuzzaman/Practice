package main

import "fmt"

func square(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
func main() {
	fmt.Println("Main is Started.....")
	c := make(chan int, 3)

	go square(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	var v int
	fmt.Scan(&v)
	fmt.Println("Program Stopped,")

}
