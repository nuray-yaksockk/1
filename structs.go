package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Привет, меня зовут %s, мне %d лет", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("С днем рождения, %s! Тебе теперь %d лет\n", p.Name, p.Age)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	person1 := Person{
		Name:  "Анна",
		Age:   25,
		Email: "anna@example.com",
	}

	fmt.Println("Person1:", person1)
	fmt.Println("Имя:", person1.Name)
	fmt.Println(person1.Introduce())

	person1.Birthday()
	fmt.Println("После дня рождения:", person1)

	car := struct {
		Brand string
		Model string
		Year  int
	}{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2022,
	}
	fmt.Println("Машина:", car)

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Прямоугольник: %.1f x %.1f\n", rect.Width, rect.Height)
	fmt.Printf("Площадь: %.1f\n", rect.Area())
}
