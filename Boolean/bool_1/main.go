// Boolean 1: Дано целое число A. Проверить истинность высказывания: «Число A является положительным»

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите любое число:")
	fmt.Scan(&a)

	var result bool

	result = a > 0

	fmt.Println(result)



}