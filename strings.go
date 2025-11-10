package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str1 := "Hello"
	str2 := "World"

	greeting := str1 + " " + str2
	fmt.Println("Приветствие:", greeting)

	russian := "Привет"
	fmt.Printf("Длина '%s' в байтах: %d\n", russian, len(russian))
	fmt.Printf("Длина '%s' в символах: %d\n", russian, utf8.RuneCountInString(russian))

	fullName := "Иван Иванов"
	firstName := fullName[:4] // "Иван"
	fmt.Println("Имя:", firstName)

	text := "Go is an open source programming language"
	fmt.Println("Original:", text)
	fmt.Println("Upper:", strings.ToUpper(text))
	fmt.Println("Lower:", strings.ToLower(text))
	fmt.Println("Contains 'open':", strings.Contains(text, "open"))
	fmt.Println("Index of 'source':", strings.Index(text, "source"))
	fmt.Println("Replace:", strings.Replace(text, "Go", "Golang", 1))

	words := strings.Split(text, " ")
	fmt.Println("Слова:", words)
	fmt.Println("Количество слов:", len(words))

	joined := strings.Join(words, "-")
	fmt.Println("Объединенные слова:", joined)
}
