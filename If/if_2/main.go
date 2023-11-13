package main

import "fmt"

func main() {
	var a int

	fmt.Println("Введите число:")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Ваше число положительно, значит ответ:",a+1)
	} else {
		fmt.Println("Ваше число отрицательно, значит ответ:",a-2)
	}
}