package main

import "fmt"

// Node is node
type Node struct {
	width, lenth int
}

func main() {
	temp := Node{10, 15}
	fmt.Println(temp)

	temp1 := Node{4, 66}

	fmt.Println(temp1, temp)

	temp1 = temp
	fmt.Println(temp1)
	//temp := Node{200, 100}  Not possible to override because it not mutable
	t1 := Node{55, 33}
	fmt.Println(t1)
}
