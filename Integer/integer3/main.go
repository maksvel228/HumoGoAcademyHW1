// Integer3: Дан размер файла в байтах. Используя операцию деления нацело,
// найти количество полных килобайтов, которые занимает данный файл (1 килобайт = 1024 байта)

package main

import "fmt"

func main() {

	var bt int

	fmt.Println("Введите количество байтов:")
	fmt.Scan(&bt)

	fmt.Println("Количество полных килобайтов:", bt / 1024)

}
