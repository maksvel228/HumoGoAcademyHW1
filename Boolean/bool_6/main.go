package main

import "fmt"

func main() {

	var a, b, c int
	var result bool

	fmt.Println("Введите три числа:")
	fmt.Scan(&a,&b,&c)

	result = a < b && b < c
	fmt.Println("Истинность высказывания:",result)

}