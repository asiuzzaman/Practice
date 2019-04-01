package main

import "fmt"

func main() {
	var sli []int
	sli = append(sli, 1)
	sli = append(sli, 5)
	var a []int
	a = append(a, sli...)
	a = append(a, 55)
	a = append(a, 44)
	res := make([]int, len(a))
	var b []int
	b = append(b, 6)
	b = append(b, 12)
	b = append(b, 18)

	copy(res, a)
	fmt.Println(res)
	res = append(res[:2], append(b, res[2:]...)...)
	fmt.Println(res)

}
