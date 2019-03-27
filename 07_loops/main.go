package main

import "fmt"

func fizzbuzz(num int) {
	for i := 1; i <= num; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	// Fizzbuzz
	fizzbuzz(20)
}
