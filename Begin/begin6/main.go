// Begin6: Даны длины рёбер a, b, с прямоугольного параллелепипеда. Найти его объём V = a * b * c
// и площадь поверхности S = 2 * (a*b + b*c + a*c)

package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Println("Введите стороны a, b, c используя пробел:")
	fmt.Scan(&a, &b, &c)

	fmt.Println("Ваш объём V =", a * b * c)

	fmt.Println("Ваша плошадь S =", 2 * (a*b + b*c + a*c))

}
