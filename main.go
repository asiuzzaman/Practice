package main

import (
	"fmt"
	"net/http"
	"path"
)

var database = make(map[interface{}]interface{})

func main() {
	fmt.Println("Ready.....")
	http.HandleFunc("/", index)
	http.HandleFunc("/set/", set)
	http.HandleFunc("/get/", get)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	for k, v := range database {

		fmt.Fprintln(w, k, v)
	}
}

func set(w http.ResponseWriter, r *http.Request) {

	fmt.Println("set korsi")
	temp := r.URL.Query()
	k := temp.Get("key")
	v := temp.Get("value")
	if k == "" || v == "" {
		fmt.Fprintln(w, "Please enter key and value correctly..")
		return
	}
	database[k] = v
	fmt.Fprintln(w, "value inserted")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getfunc")
	temp := path.Base(r.URL.Path)
	fmt.Fprintln(w, database[temp])

}
