package main

import (
	"fmt"
)

func main() {

	// Declaring array with size
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// Declaring array without size
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{9, 8, 7, 6}

	fmt.Println(arr3)
	fmt.Println(arr4)

	//Declaring String
	var cars = [...]string{"BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	// Array access elements
	prices := [3]int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Initialize only specific elements
	arr5 := [5]int{1: 10, 2: 40}
	fmt.Println(arr5)

	// Find the lenght of an array using len()
	fmt.Println(len(arr5))
}
