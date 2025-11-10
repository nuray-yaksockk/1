package main

import "fmt"

func main() {

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	var names [3]string

	names[0] = "Алиса"
	names[1] = "Боб"
	names[2] = "Чарли"

	fmt.Println("Массив numbers:", numbers)
	fmt.Println("Массив names:", names)

	// Перебор массива
	fmt.Println("Элементы numbers:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}

	fmt.Println("Элементы names:")
	for index, value := range names {
		fmt.Printf("names[%d] = %s\n", index, value)
	}

	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Матрица 2x3:", matrix)
}
