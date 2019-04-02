package main

import "fmt"

type Contact struct {
	phone, address string
}
type Employee struct {
	FirstName, LastName string
	contact             Contact
}

// Method on nested Structure
func (em *Employee) changePhone(newPhone string) {
	em.contact.phone = newPhone
}
func main() {
	e := Employee{
		FirstName: "Mohammad",
		LastName:  "Asiuzzaman",
		contact:   Contact{"01725955624", "ItechSoftSolution,Khulna,Bangladesh"},
	}

	fmt.Println(e)
	e.changePhone("01558951384")
	fmt.Println(e)
}
