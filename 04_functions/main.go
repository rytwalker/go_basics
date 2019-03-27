package main

import "fmt"

// function name (param type) func type
func greeting(name string) string {
	return "Hello " + name
}

// func getSum(num1, num2 int) if same type
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Ryan"))
	fmt.Println(getSum(50, 473))
}
