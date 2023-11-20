package main

import "fmt"

func main() {

	var a float64

	fmt.Println("Введите цену за кг:")
	fmt.Scan(&a)

	for i := 12; i <= 20; i += 2 {

		result := float64(i) / 10
		result1 := a * float64(i)

		fmt.Println("Килограммы:", result, "Цена:", result1)

	}

}
