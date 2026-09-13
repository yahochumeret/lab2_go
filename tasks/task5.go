package tasks

import "fmt"

// Rectangle описывает прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

// Area вычисляет площадь прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter вычисляет периметр (бонус)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Task5 демонстрирует работу со структурой Rectangle
func Task5() {
	rect := Rectangle{Width: 10.5, Height: 4.2}

	fmt.Printf("Прямоугольник: ширина = %.2f, высота = %.2f\n",
		rect.Width, rect.Height)
	fmt.Printf("Площадь: %.2f\n", rect.Area())
	fmt.Printf("Периметр: %.2f\n", rect.Perimeter())

	// Второй прямоугольник
	rect2 := Rectangle{Width: 3, Height: 7}
	fmt.Printf("\nВторой прямоугольник: %.2f x %.2f\n",
		rect2.Width, rect2.Height)
	fmt.Printf("Площадь: %.2f\n", rect2.Area())
}
