package main

import "fmt"

func main() {

	var a, b int
	var result bool

	fmt.Println("Введите два целых числа:")
	fmt.Scan(&a, &b)

	result = a%2 == 1 && b%2 == 1

	fmt.Println("Истинноть высказывания:", result)

}
