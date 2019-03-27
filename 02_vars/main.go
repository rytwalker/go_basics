package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8

	// using var
	var name = "Ryan"
	var age = 29

	// shorthand variable
	animal := "Bird"
	// very shorthand
	email, password := "rytwalker@gmail.com", "password"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(animal)
	fmt.Println(email)
	fmt.Println(password)
	fmt.Printf("%T\n", animal)
}
