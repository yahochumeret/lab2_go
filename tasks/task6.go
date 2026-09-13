package tasks

import "fmt"

// AverageInt возвращает среднее значение двух целых чисел как float64
func AverageInt(a, b int) float64 {
	return float64(a+b) / 2.0
}

// AverageIntExact возвращает среднее значение с округлением вниз (int)
func AverageIntExact(a, b int) int {
	return (a + b) / 2
}

// Task6 демонстрирует вычисление среднего значения
func Task6() {
	pairs := [][2]int{
		{10, 20},
		{7, 8},
		{15, 22},
		{-5, 5},
	}

	for _, p := range pairs {
		a, b := p[0], p[1]
		fmt.Printf("Среднее(%d, %d) = %.1f (точно: %d)\n",
			a, b, AverageInt(a, b), AverageIntExact(a, b))
	}
}
