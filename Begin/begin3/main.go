// Begin3: Даны стороны прямоугольника a и b. Найти его площадь S = a * b и периметр P = 2(a + b)

package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Введите сторону a:")
	fmt.Scan(&a)

	fmt.Println("Введите сторону b:")
	fmt.Scan(&b)

	fmt.Println("Ваша площадь S =", a*b)

	fmt.Println("Ваш периметр P =", 2*(a+b))

}
