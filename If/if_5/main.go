package main

import "fmt"

func main() {
	var a, b, c, result, result1 int

	fmt.Println("Введите три целых числа:")
	fmt.Scan(&a, &b, &c)

	result, result1 = 0, 0

	if a > 0 {
		result++
	} else {
		result1++
	}
	if b > 0 {
		result++
	} else {
		result1++
	}
	if c > 0 {
		result++
	} else {
		result1++
	}
	fmt.Println("Количество положительных чисел:", result)
	fmt.Println("Количество отрицательных чисел:", result1)

}
