// Boolean 5: Даны два целых числа: A, B. Проверить истинность высказывания: «Справедливы неравенства A ≥ 0 или B < −2

package main

import "fmt"

func main() {
	
	var a,b int
	var result bool

	fmt.Println("Введите два числа через пробел:")
	fmt.Scan(&a, &b)

	result = a >= 0 && b < -2

	fmt.Println(result)


}