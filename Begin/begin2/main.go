// begin2: Дана сторона квадрата a. Найти его площадь S=a*a

package main

import (
	"fmt"
)

func main() {

	var a int

	fmt.Println("Введите сторону квадрата a:")

	fmt.Scan(&a)

	fmt.Println("Ваша площадь S равна:", a*a)

}
