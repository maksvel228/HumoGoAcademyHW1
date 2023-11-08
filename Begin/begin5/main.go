// Begin5: Дана длина ребра куба a. Найти объём куба V = a * a * a и площадь его поверхности S = 6 * a * a

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите длину ребра куба a:")
	fmt.Scan(&a)

	fmt.Println("Ваш объём куба V =", a*a*a)

	fmt.Println("Ваша площадь поверхности S =", 6*(a*a))

}
