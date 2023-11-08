// Begin4: Дан диаметр окружности d. Найти её длину L = PI * d. В качестве значения PI использовать 3.14

package main

import "fmt"

const PI = 3.14

func main() {

	var d float32

	fmt.Println("Введите окружость d:")
	fmt.Scan(&d)

	fmt.Println("Ваша длина L =", PI * d)

}
