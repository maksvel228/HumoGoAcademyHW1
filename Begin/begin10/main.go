// Begin10: Даны два ненулевых числа. Найти сумму, разность, произведение и частное их квадратов

package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64

	fmt.Println("Введите значения a и b через запятую")
	fmt.Scan(&a, &b)

	fmt.Println("Сумма квадратов a и b:", math.Sqrt(a + b))
	fmt.Println("Разность квадратов a и b:", math.Sqrt(a - b))
	fmt.Println("Произведение квадратов a и b:", math.Sqrt(a * b))
	fmt.Println("Частное квадратов a и b:", math.Sqrt(a / b))

}
