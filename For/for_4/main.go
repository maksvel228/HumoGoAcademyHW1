package main

import "fmt"

func main() {
	
	var a float64

	fmt.Println("Цена за 1 кг:")
	fmt.Scan(&a)

	for i := 1; i <= 10; i++ {

		result := a * float64(i)

		fmt.Println("Килограммы:",i,"Цена:",result)

	}



}