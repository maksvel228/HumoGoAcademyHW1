// Integer6: Дано двузначное число. Вывести вначале его левую цифру (десятки), а затем — его правую цифру (единицы).
// Для нахождения десятков использовать операцию деления нацело, для нахождения единиц — операцию взятия остатка от деления

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите двухзначное число:")
	fmt.Scan(&a)

	fmt.Println("Левая цифра вашего числа:", a / 10, "| Правая цифра вашего числа:", a % 10)



}
