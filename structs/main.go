package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Rogachev",
		contact: contactInfo{
			email:   "aleksandr.rogachev1994@gmail.com",
			zipCode: 11111,
		},
	}

	// alexPt := &alex
	alex.updateName("Alex2")
	alex.print()
	fmt.Printf("%p\n", &alex)
}

func (pPerson *person) updateName(newFirstName string) {
	(*pPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
