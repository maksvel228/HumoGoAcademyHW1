package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Two number A < B:")
	fmt.Scan(&a, &b)

	result := 0

	for i := a; i <= b; i++ {

		fmt.Println(i)
		result++

	}
	fmt.Println("Целых чисел:",result)
	



}
