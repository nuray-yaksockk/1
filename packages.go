package main

import (
	"fmt"
	"homework/mathutil"
	"strings"
)

func main() {
	sum := mathutil.Add(10, 5)
	product := mathutil.Multiply(4, 7)

	fmt.Printf("10 + 5 = %d\n", sum)
	fmt.Printf("4 * 7 = %d\n", product)
	fmt.Printf("Число 10 четное? %t\n", mathutil.CheckEven(10))

	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.Repeat("-", 20))

}
