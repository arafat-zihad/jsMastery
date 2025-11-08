package main

import (
	"fmt"
)

func main() {
	slice1 := []int{}
	fmt.Println(len(slice1))

	fmt.Println(cap(slice1))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// Creating Slice from an Array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))

	// Create a slice with The make()
	slice2 := make([]int, 5, 10)
	fmt.Printf("Slice 2 = %v\n", slice2)
	fmt.Printf("Length = %d\n", len(slice2))
	fmt.Printf("Capacity = %d\n", cap(slice2))

	// Add element using append()
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	arr2 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice3 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	myslice3 = arr2[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	myslice3 = append(myslice3, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

}
