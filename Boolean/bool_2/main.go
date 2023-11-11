// Boolean 2: Дано целое число A. Проверить истинность высказывания: «Число A является нечетным»

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите любое число:")
	fmt.Scan(&a)

	var result bool

	result = a%2 == 1

	fmt.Println(result)

}
