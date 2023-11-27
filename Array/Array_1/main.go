package main

import "fmt"

func main() {

	var n int

	fmt.Println("Размер массива:")
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var x int
	fmt.Println("Введите число X:")
	fmt.Scan(&x)

	res := 0

	for i := 0; i < n; i++ {
		if a[i] == x {
			res++
		}
	}
	fmt.Println("Число X встречается:", res, "раз")

}
