package main

import "fmt"

func main() {

	number := 2

	fmt.Printf("Проверяем число: %d\n", number)

	// Проверяем условия
	if number == 0 {
		fmt.Println("Число является нулем")
	} else if number > 0 {
		if number%2 == 0 {
			fmt.Printf("Число %d - положительное четное\n", number)
		} else {
			fmt.Printf("Число %d - положительное нечетное\n", number)
		}
	} else {
		fmt.Printf("Число %d - отрицательное\n", number)
	}
}
