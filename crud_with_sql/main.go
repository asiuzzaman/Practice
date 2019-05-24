package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


type employee struct {
	ID         int
	first_name string
	last_name  string
}

func main() {
	fmt.Println("My crud application...........")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	db.Query("drop table employee")
	_, err = db.Query("create table employee(id int not null auto_increment, first_name varchar(50), last_name varchar(30), primary key(id));")

	if err != nil {
		panic(err.Error())
	}

	db.Query("insert into employee (first_name,last_name) values('mohammad','asiuzzaman')")
	db.Query("insert into employee (first_name,last_name) values('farat','zim')")

	read, err := db.Query("select * from employee")
	if err != nil {
		panic(err.Error())
	}
	for read.Next() {
		var res employee
		err = read.Scan(&res.ID, &res.first_name, &res.last_name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(res)
	}

	db.Query("insert into employee (first_name,last_name) values('Itech','khulna')")

	read, err = db.Query("select * from employee")
	if err != nil {
		panic(err.Error())
	}

	for read.Next() {
		var res employee
		err = read.Scan(&res.ID, &res.first_name, &res.last_name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(res)
	}
	fmt.Println("After Update..........")
	db.Query("update employee set first_name='bangladesh', last_name='dhaka' where ID=3")

	read, err = db.Query("select * from employee")
	if err != nil {
		panic(err.Error())
	}

	for read.Next() {
		var res employee
		err = read.Scan(&res.ID, &res.first_name, &res.last_name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(res)
	}

	fmt.Println("Deleting..............................")
	db.Query("delete from employee where id=3")

	read, err = db.Query("select * from employee")
	if err != nil {
		panic(err.Error())
	}

	for read.Next() {
		var res employee
		err = read.Scan(&res.ID, &res.first_name, &res.last_name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(res)
	}

	// result, err := db.Query("SELECT * FROM Books")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// for result.Next() {
	// 	var user Books
	// 	err = result.Scan(&user.BookName, &user.PageNo)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(user)

	// }

	// //	db.Query("update customers set NAME='bangladesh' where ID=3")

	// // fmt.Println(insert)
	// //defer insert.Close()
	// // fmt.Println("After Updating the Name.....")
	// // result, err = db.Query("SELECT * FROM customers")

	// // if err != nil {
	// // 	panic(err.Error())
	// // }
	// // for result.Next() {
	// // 	var user customers
	// // 	err = result.Scan(&user.ID, &user.NAME, &user.age, &user.salary)
	// // 	if err != nil {
	// // 		panic(err.Error())
	// // 	}
	// // 	fmt.Println(user)

	// // }

	fmt.Println("successfully created.............")

	defer db.Close()
}
