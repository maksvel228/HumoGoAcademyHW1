package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Two number A < B:")
	fmt.Scan(&a, &b)

	result := 1

	for i := a; i <= b; i++ {

		fmt.Println("Ваши числа:", i)
		result *= i

	}
	fmt.Println("Произведение чисел:", result)
}
