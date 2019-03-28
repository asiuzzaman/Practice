package main

import "fmt"

func main() {
	arr := []float64{4, 5, 76, 7, 11, 90}
	average := avg(arr)
	fmt.Println(arr, average)
	anoarr := []float64{4, 6, 9, 0}
	pass(anoarr...)
	r := passone(anoarr...)
	fmt.Println(r)
}

func avg(name []float64) float64 {

	var s float64 = 0
	for _, v := range name {
		s = s + v
	}
	res := s / float64(len(name))
	return res
}

func pass(naming ...float64) {
	fmt.Println(naming)
}

func passone(arr ...float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
