package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите число:")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Ваш конечный ответ:",a+1)
	} else {
		fmt.Println("Ваш конечный ответ:",a)
	}

}