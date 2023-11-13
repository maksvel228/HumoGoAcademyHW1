package main

import "fmt"

func main() {
	var a, b, c, result int

	fmt.Println("Введите три числа через пробел:")
	fmt.Scan(&a, &b, &c)

	result = 0

	if a > 0 {
		result++
	}
	if b > 0 {
		result++

	}
	if c > 0 {
		result++
	}
	fmt.Println("Количество положительных чисел:", result)

}
