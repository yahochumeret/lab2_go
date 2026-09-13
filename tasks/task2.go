package tasks

import "fmt"

// CheckSign возвращает "Positive", "Negative" или "Zero"
func CheckSign(n int) string {
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	}
	return "Zero"
}

// Task2 демонстрирует работу функции CheckSign
func Task2() {
	numbers := []int{5, -3, 0, 42, -17}

	for _, n := range numbers {
		fmt.Printf("Число %d → %s\n", n, CheckSign(n))
	}
}
