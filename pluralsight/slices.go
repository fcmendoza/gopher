package main

import (
	"fmt"
)

func main() {
	// Declares a slice called myCourses
	//myCourses := make([]string, 5, 10)
	// myCourses := []string { "Docker", "Puppet", "Python" }

	// fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(myCourses), cap(myCourses))

	// mySlice := []int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	// fmt.Println(mySlice[4])

	// mySlice[1] = 0
	// fmt.Println(mySlice)

	// sliceOfSlice := mySlice[2:5]
	// fmt.Println(sliceOfSlice)

	// mySlice := make([]int, 1, 4)

	// fmt.Printf("Length is: %d. \nCapacity is: %d\n", 
	// 	len(mySlice), cap(mySlice))

	// for i := 1; i < 17; i++ {
	// 	mySlice = append(mySlice, i)
	// 	fmt.Printf("\nCapacity is: %d", cap(mySlice))
	// }

	mySlice := []int {1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	for _, i := range mySlice {
		fmt.Println("for range loop:", i)
	}

	newSlice := []int {10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}

