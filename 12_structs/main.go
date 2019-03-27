package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// greeting method (value reciever)
func (person Person) greet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer reciever)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer reciver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		person.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Gary", lastName: "Fodge", city: "Boston", gender: "m", age: 24}
	// person1 := Person{"Gary", "Fodge", "Boston", "m", 24}

	fmt.Println(person1)

	// single field
	fmt.Println(person1.firstName)
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Barrison")
	fmt.Println(person1.greet())
}
