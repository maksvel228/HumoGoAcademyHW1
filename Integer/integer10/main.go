// Integer10: Дано трехзначное число. Вывести вначале его последнюю цифру(единицы), а затем — его среднюю цифру (десятки)

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите трёхзначное число:")
	fmt.Scan(&a)

	fmt.Println("Последняя цифра вашего числа:", a%10, "\nСредняя цифра вашего числа:", a/10%10)

}
