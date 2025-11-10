package main

import "fmt"

func main() {

	studentGrades := map[string]int{
		"Алиса": 95,
		"Боб":   87,
		"Чарли": 92,
	}

	fmt.Println("Оценки студентов:", studentGrades)

	studentGrades["Диана"] = 88
	fmt.Println("После добавления Дианы:", studentGrades)

	grade, exists := studentGrades["Алиса"]
	if exists {
		fmt.Printf("Оценка Алисы: %d\n", grade)
	}

	delete(studentGrades, "Боб")
	fmt.Println("После удаления Боба:", studentGrades)

	fmt.Println("Все студенты и оценки:")
	for name, grade := range studentGrades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	ages := make(map[string]int)
	ages["Иван"] = 25
	ages["Мария"] = 30
	fmt.Println("Возрасты:", ages)
}
