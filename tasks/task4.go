package tasks

import (
	"fmt"
	"unicode/utf8"
)

// StringLength возвращает длину строки в символах (рунах)
func StringLength(s string) int {
	return utf8.RuneCountInString(s)
}

// StringLengthBytes возвращает длину строки в байтах
func StringLengthBytes(s string) int {
	return len(s)
}

// Task4 демонстрирует работу с длиной строки
func Task4() {
	examples := []string{"Hello", "Привет", "Go", "世界"}

	for _, s := range examples {
		fmt.Printf("Строка: %-8q | Символов: %d | Байт: %d\n",
			s, StringLength(s), StringLengthBytes(s))
	}
}
