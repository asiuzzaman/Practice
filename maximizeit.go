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
		//fmt.Println(niwi[i].ni, " ", niwi[i].wi)

	}

	mysort(niwi)

	var temp = 0
	for i := 0; i < N-1; i++ {
		temp = temp + niwi[i].wi
		if temp <= K {
			fmt.Print(niwi[i].ni, ",")
		}
	}
	fmt.Println(niwi[N-1].ni)

}
