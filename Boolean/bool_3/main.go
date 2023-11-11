// Boolean 3: Дано целое число A. Проверить истинность высказывания: «Число A является четным»

package main

import "fmt"

func main() {

	var a int

	var result bool

	fmt.Println("Введите любое число:")
	fmt.Scan(&a)

	result = a%2 == 0

	fmt.Println(result)

}
