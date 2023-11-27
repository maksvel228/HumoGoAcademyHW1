package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Длина массива:")
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var x int
	fmt.Println("Введите число X:")
	fmt.Scan(&x)

	res := a[0]
	res1 := x - res


	
}
