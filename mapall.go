package main

import (
	"fmt"
	"sort"
)

func main() {

	mp := make(map[string]int)
	mp["asi"] = 56
	mp["uzz"] = 11
	mp["aman"] = 66
	mp["bang"] = 71
	mp["desh"] = 52

	delete(mp, "aman")

	fmt.Println(mp)

	for key, value := range mp {
		fmt.Println(key, value)
	}

	var temp []string

	for k := range mp {
		temp = append(temp, k)
	}

	sort.Strings(temp)

	for _, value := range temp {
		fmt.Println("here: ", mp[value])
	}

	//fmt.Println(temp)

	comMap := map[string]map[string]string{
		"man": map[string]string{
			"asiu":  "zzaman",
			"level": "geeks",
		},
		"woman": map[string]string{
			"my":     "Girlfriend",
			"level0": "silly",
		},
	}

	fmt.Print(comMap)
}
