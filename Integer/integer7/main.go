// Integer7: Дано двузначное число. Найти сумму и произведение его цифр

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите двухзначное число:")
	fmt.Scan(&a)

	fmt.Println("Для начала мы выведим правую и левую цифру вашего числа\nЛевая:", a/10, "| Правая:", a%10)

	var result1 = a / 10 + a % 10
	fmt.Println("Сумма ваших цифр:", result1)

	var result2 = (a / 10) * (a % 10)
	fmt.Println("Произведение ваших цифр:", result2)

}
