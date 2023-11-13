package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите два числа через пробел:")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Порядковый номер наименьшего числа:", 2)
	} else {
		fmt.Println("Порядковый номер наименьшего числа:", 1)
	}
}
