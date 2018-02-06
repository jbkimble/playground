package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string
	zipcode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo {
			email: "j@gmail.com",
			zipcode: 94000,
		},
	}

// a reference to the memory address that the jim struct exists at
	jim.updateName("Panda")
	jim.print()
}

func (p person) print() {
	fmt.Println("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
