// Integer1: Дано расстояние L в сантиметрах
// Используя операцию деления нацело, найти количество полных метров в нем (1 метр = 100 см)

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите расстояние:")
	fmt.Scan(&a)

	fmt.Println("Количество полных метров:", a / 100)

}
