// Boolean 4: Даны два целых числа: A, B. Проверить истинность высказывания: «Справедливы неравенства A > 2 и B ≤ 3»

package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Введите два числа через пробел:")
	fmt.Scan(&a, &b)

	var result bool

	result = a > 2 && b <= 3

	fmt.Println(result)

}
