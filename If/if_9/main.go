package main

import "fmt"

func main() {

	var a, b float32

	fmt.Println("Введите два числа через пробел, значение A должны быть больше:")
	fmt.Scan(&a, &b)

	fmt.Println("Значение числа A:", a, "| Значение числа B:", b)

	if a > b {

		b = a
		fmt.Println("Новое значение числа B:", b)

		a = b
		fmt.Println("Новое значение числа A:", a)

	} else {
		fmt.Println("Условие задачи не соблюдено!")
	}

}
