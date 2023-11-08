// Integer2: Дана масса M в килограммах, используя операцию деления нацело, найти количество полных
// тонн в ней (1 тонна = 1000 кг)

package main

import "fmt"

func main() {

	var M int

	fmt.Println("Введите массу в кг.:")
	fmt.Scan(&M)

	fmt.Println("Ваша масса полных тонн:", M / 1000)

}
