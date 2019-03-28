package main

import "fmt"

func main() {
	var sli []int
	sli = append(sli, 5)

	for i := 0; i < 5; i++ {
		sli = append(sli, i*3)
		fmt.Println(sli)
	}
	sli[3] = 777 /// insert the value in a specific position.
	fmt.Println(sli)

}
