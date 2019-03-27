package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	// assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Danny"] = "danny@gmail.com"
	// emails["Caren"] = "caren@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@gmai.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// delete
	delete(emails, "Bob")
	fmt.Println(emails)
}
