package main

import "fmt"

func main() {

	var numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", numbers)

	numbers = append(numbers, 6)
	numbers = append(numbers, 7, 8, 9)
	fmt.Println("После добавления:", numbers)

	slice1 := numbers[2:5]
	slice2 := numbers[:3]
	slice3 := numbers[6:]

	fmt.Println("numbers[2:5]:", slice1)
	fmt.Println("numbers[:3]:", slice2)
	fmt.Println("numbers[6:]:", slice3)

	dynamicSlice := make([]int, 3, 5)
	dynamicSlice[0] = 10
	dynamicSlice[1] = 20
	dynamicSlice[2] = 30
	fmt.Println("Динамический слайс:", dynamicSlice)
	fmt.Printf("Длина: %d, Емкость: %d\n", len(dynamicSlice), cap(dynamicSlice))

	copySlice := make([]int, len(numbers))
	copy(copySlice, numbers)
	fmt.Println("Копия numbers:", copySlice)
}
