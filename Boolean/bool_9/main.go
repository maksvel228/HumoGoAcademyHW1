package main

import "fmt"

func main() {

	var a, b int
	var result bool

	fmt.Println("Введите два целых числа:")
	fmt.Scan(&a, &b)

	result = (a%2 == 1 && b%2 == 1) || (a%2 == 0 && b%2 == 0)

	fmt.Println("Истинность высказывания:", result)

}
