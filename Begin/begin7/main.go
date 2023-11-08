// Begin7: Найти длину окружности L и площадь круга S заданного радиуса R: L = 2 * PI * R, S = PI * R * R

package main

import "fmt"

const PI = 3.14

func main() {

	var r float32

	fmt.Println("Задаём значение R:")
	fmt.Scan(&r)

	fmt.Println("Ваша длина окружности L:", 2 * PI * r)

	fmt.Println("Ваша площадь круга S:", PI * (r * r))

}
