// Begin8: Даны два числа a и b. Найти их среднее арифмитическое (a + b) / 2

package main

import "fmt"

func main() {

	var a, b float32

	fmt.Println("Вводим значения a и b через запятую:")
	fmt.Scan(&a, &b)

	fmt.Println("Ваше среднее арифмитическое:", (a + b) / 2)

}
