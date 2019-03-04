package main

import (
	"fmt"
	"reflect"
)

// var (
// 	name, course string  // name of subscriber and current course
// 	module       float64 // current place in course
// )

// var (
// 	name, coure, module = "John", "Docker Deep Dive", 3.2
// )

// var (
// 	name   = "Jon"              // Name of subscriber∏
// 	course = "Docker Deep Dive" // Course being viewed
// 	module = 3.6                // Current position in course
// )

func main() {
	name := "Jon" // Name of subscriber∏
	// course := "Docker Deep Dive" // Course being viewed
	module := 3.6 // Current position in course

	fmt.Println("Name is set to", name, "and is fo type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is fo type", reflect.TypeOf(module))

	// a := 10.0
	// b := 3
	// c := int(a) + b
	// fmt.Println("C has value:", c, "and is fo type:", reflect.TypeOf(c))
}
