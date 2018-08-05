package main

import (
	"fmt"
//	"reflect"
	"os"
)

// var (
// 	name   string  // Name of subscriber
// 	course string  // Name of current course
// 	module float64 // Current place in course
// )

// var (
// 	name, course string
// 	module       float64
// )

// var (
// 	name, course, module = "Nigel", "Docker Deep Dive", 3.2
// )

// var (
// 	name   = "Nigel"            // Name of subscriber
// 	course = "Docker Deep Dive" // Course being viewed
// 	module = 3.2                // Current position in course
// )

func main() {
	//name   := "Nigel"            // Name of subscriber
	name   := os.Getenv("USER")            // Name of subscriber
	course := "Docker Deep Dive" // Course being viewed
	//module := 3.2                // Current position in course
	//ptr := &module

	// fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	// fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	// //fmt.Println("Memory address of *module* variable is", &module)
	// fmt.Println("Memory address of *module* variable is", 
	// 		ptr, "and the value if *module* is", *ptr)

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYour're now watching course", course)

	const speedOfLight = 180000

	fmt.Println("Speed of light is", speedOfLight)
}

func changeCourse(course *string) string {
	*course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}



