package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Println("Welcome to the GO course, ", name)
}

func printSomething() {
	fmt.Println("Now print the addition, multiplication")
}

func add(n1 int, n2 int) int {
	sum := n1 + n2

	return sum
}

func getNumbers(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	mul := n1 * n2

	return sum, mul
}

func main() {
	a := 10
	b := 20

	sayHello("Zihad")
	printSomething()

	sum := add(a, b)

	p, q := getNumbers(a, b)

	fmt.Println(p)
	fmt.Println(q)

	fmt.Println(sum)
}
