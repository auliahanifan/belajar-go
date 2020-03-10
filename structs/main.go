package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// when initialize variable but doesnt refer any value, it become Zero Value
	// var budi person

	// budi.firstName = "Budi"
	// budi.lastName = "Sudarman"

	alex := person{
		firstName: "Alex",
		lastName:  "Firmansyah",
		contact: contactInfo{
			email: "email@email.com",
			zip:   057,
		},
	}

	alex.updateFirstName("JOSHUA")
	alex.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
