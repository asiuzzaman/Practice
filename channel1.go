package main

import (
	"fmt"
	"time"
)

func send(ch chan string, str string) {
	ch <- str
}

func receive(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	var ch = make(chan string)
	go receive(ch)

	send(ch, "hello1")
	send(ch, "hello2")
	send(ch, "hello3")

	time.Sleep(time.Second)
}
