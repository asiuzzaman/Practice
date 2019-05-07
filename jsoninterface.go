package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonData := []byte(`{"Name":"asiu","age":22,"parents":["myfather","mymother"]}`)
	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})
	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("  ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}
