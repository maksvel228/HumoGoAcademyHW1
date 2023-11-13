package main

import "fmt"

func main() {

	var a, b, result int

	fmt.Println("Введите два числа через пробел:")
	fmt.Scan(&a, &b)

	if a != b {

		result = a + b
		fmt.Println("Ваши значения не равны, новые значения ваших переменных:")
		fmt.Println("A:", a+result, "| B:", b+result)

	} else if a == b {
		a = 0
		b = 0
		fmt.Println("Ваши значения равны, новые значения ваших переменных:")
		fmt.Println("A:",a,"| B:", b)

	}

}
