package main

import "fmt"

func main() {

	x := 10
	var p *int = &x

	fmt.Printf("x = %d\n", x)
	fmt.Printf("Адрес x: %p\n", &x)
	fmt.Printf("p = %p\n", p)
	fmt.Printf("*p = %d\n", *p)

	*p = 20
	fmt.Printf("После *p = 20: x = %d\n", x)

	y := 5
	fmt.Printf("До функции: y = %d\n", y)
	double(&y)
	fmt.Printf("После функции: y = %d\n", y)

	var pp **int = &p
	fmt.Printf("pp = %p\n", pp)
	fmt.Printf("*pp = %p\n", *pp)
	fmt.Printf("**pp = %d\n", **pp)

	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Printf("Первый элемент массива: %d\n", (*arrPtr)[0])
}

func double(ptr *int) {
	*ptr = *ptr * 2
	fmt.Printf("В функции: значение = %d\n", *ptr)
}
