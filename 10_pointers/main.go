package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Println(a, *b)

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
