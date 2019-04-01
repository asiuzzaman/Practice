package main

import (
	"errors"
	"fmt"
)

type student struct {
	name string
	age  int
}

func addStudent(s student) error {

	db := make(map[string]student)

	if s.age > 30 {
		return errors.New("Student is too Old")

	}
	db[s.name] = s
	fmt.Println(db)
	return nil
}

func main() {
	tom := student{name: "Tom", age: 45}
	//jerry := student{name: "Jerry", age: 10}

	err := addStudent(tom)

	fmt.Println("we got an err " + err.Error())

	//addStudent(jerry)

	//fmt.Println(db["Tom"])
	// err := addStudent(tom)
	// fmt.Println("We got an error " + err.Error())

}
