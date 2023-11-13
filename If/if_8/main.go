package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Введите два числа через пробел:")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Наибольшее:", a, "| Наименьшее:", b)
	} else {
		fmt.Println("Наибольшее:", b, "| Наименьшее:", a)
	}

}
