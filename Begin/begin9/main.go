// Begin9: Даны два неотрицательных числа a и b. Найти их среднее геометрическое,
// то есть квадратный корень из их произведения: a * b

package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64

	fmt.Println("Введите через запятую значения a и b:")
	fmt.Scan(&a, &b)

	fmt.Println("Ваше среднее геометрическое значение:", math.Sqrt(a * b))

}
