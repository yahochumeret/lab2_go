package tasks

import "fmt"

// Task1 определяет, является ли введённое число чётным или нечётным
func Task1() {
	var number int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("Число %d — чётное\n", number)
	} else {
		fmt.Printf("Число %d — нечётное\n", number)
	}
}
