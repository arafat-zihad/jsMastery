package main

import (
	"fmt"
)

func main() {

	// Variable Declaration with inicial Value
	var std1 string = "Jhon"
	var std2 = "Zihad"
	x := 2

	fmt.Println(std1)
	fmt.Println(std2)
	fmt.Println(x)

	// Variable Declaration without inicial Value
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Multiple Variable Declaration
	var a1, b1 = 6, "Hello"
	c1, d1 := 7, "World"

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

	// Variable Declaraton in block
	var (
		a2 int
		b2 int    = 1
		c2 string = "Hola"
	)

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

}
