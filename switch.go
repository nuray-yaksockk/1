package main

import "fmt"

func main() {
	day := 3

	fmt.Printf("День недели: %d\n", day)

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6, 7:
		fmt.Println("Выходной!")
	default:
		fmt.Println("Неверный день недели")
	}

	// Switch без выражения
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Оценка: A")
	case score >= 80:
		fmt.Println("Оценка: B")
	case score >= 70:
		fmt.Println("Оценка: C")
	default:
		fmt.Println("Оценка: F")
	}
}
