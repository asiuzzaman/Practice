package main

import (
	"fmt"
	"sort"
)

type Node struct {
	ni, wi int
}

func mysort(nodes []Node) {
	sort.SliceStable(nodes, func(i, j int) bool {
		mi, mj := nodes[i], nodes[j]

		return mi.wi < mj.wi
	})
}

func main() {
	var N, K int
	fmt.Scanln(&N)

	fmt.Scanln(&K)
	niwi := make([]Node, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&niwi[i].ni)
		fmt.Scanln(&niwi[i].wi)
	}

	mysort(niwi)

	if niwi[0].wi > K {
		fmt.Println("-1")
		return
	}

	var temp = 0
	store := []int{}

	for i := 0; i < N; i++ {
		temp = temp + niwi[i].wi
		if temp <= K {
			store = append(store, niwi[i].ni)
		} else {
			break
		}

	}
	sort.Ints(store)

	fmt.Print(store[0])
	for i := 1; i < N; i++ {
		fmt.Print(",", store[i])
	}

}
