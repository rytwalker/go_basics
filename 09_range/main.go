package main

import "fmt"

func main() {
	ids := []int{33, 43, 35, 31, 65, 44}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// no index
	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// range with map
	emails := map[string]string{"Bob": "bob@gmai.com", "Sharon": "sharon@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
