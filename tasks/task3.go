package tasks

import "fmt"

// Task3 выводит числа от 1 до 10 с помощью цикла for
func Task3() {
	fmt.Println("Числа от 1 до 10:")

	// Классический цикл for
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Альтернативный вариант — for как while
	fmt.Println("\nЧисла от 10 до 1 (обратный порядок):")
	j := 10
	for j >= 1 {
		fmt.Printf("%d ", j)
		j--
	}
	fmt.Println()
}
