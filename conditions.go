package main

import "fmt"

func main() {
	age := 19

	fmt.Printf("Возраст: %d\n", age)

	if age < 0 {
		fmt.Println("Ошибка: возраст не может быть отрицательным!")
	} else if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age == 18 {
		fmt.Println("Вам ровно 18 лет!")
	} else {
		fmt.Println("Вы совершеннолетний")
	}

	if age%2 == 0 {
		fmt.Println("Возраст четный")
	} else {
		fmt.Println("Возраст нечетный")
	}
}
