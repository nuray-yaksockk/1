package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return x / y, nil
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	sum := add(10, 5)
	fmt.Printf("10 + 5 = %d\n", sum)

	result, err := divide(15, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", result)
	}

	fmt.Printf("Факториал 5 = %d\n", factorial(5))
}
