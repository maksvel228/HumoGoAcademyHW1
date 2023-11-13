package main

import "fmt"

func main() {
	var a int

	fmt.Println("Введите число:")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("У вас положительное число, ответ:", a+1)
	} else if a < 0 {
		fmt.Println("У вас отрицательное число, ответ:", a-2)

	} else if a == 0 {
		a = 10
		fmt.Println("Ваше число равно нулю, новый ответ:", a)
	}

}
