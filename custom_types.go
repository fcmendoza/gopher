// Thanks to https://medium.com/@kandros/go-function-type-reusable-function-signatures-2389f6bdd4f6
package main

import (
	"fmt"
	"strconv"
)


/*
// typical function
func {name_of_the_function} ({params}) {return_type} {
	// function body
}
*/

func convert (plan string, cycle int) string {
	return plan + " " + strconv.Itoa(cycle)
}

// funciton as a type. we're creating/designing a new type. our type will be a custom function wuth a custom signature

//type {name_of_type} {type} {bocy}

//e.g.:

type User struct { FirstName, LastName string }

// user-defined function type
// you can declare a function signature as a type
type Converter func(plan string, cycle int) string
type greeter = func(a, b string) string
//type Int32 = implment_an_int_somehow

func NewConverter() Converter {
	return func(plan string, cycle int) string {
		return "(NewConverter)\t" + plan + " - " + strconv.Itoa(cycle)
	}
}

func NewConverter2(extra string) Converter {
	return func(plan string, cycle int) string {
		return "(NewConverter2)\t" + plan + " - " + strconv.Itoa(cycle) + ". Extra: " + extra
	}
}


func main() {
	var explicit greeter = func(a, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}

	implicit := func(a, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	printFunction(explicit)
	printFunction(implicit)
	printFunction(func(a, b string) string {
		return fmt.Sprintf("%s %s!", a, b)
	})

	//var age Int32 = 1

	var converter1 Converter = func (plan string, cycle int) string {
		return  plan + strconv.Itoa(cycle)
	}

	var converter2 = NewConverter()

	converter3 := NewConverter()

	// r := convert("Plan No.", 1)
	// fmt.Println(r)

	// c := Converter("Plan NUM.", 2)
	// fmt.Println(c)

	summaryGetter(converter1)
	summaryGetter(converter2)
	summaryGetter(converter3)
	summaryGetter(NewConverter2("yo"))
	summaryGetter(NewConverter2("hi"))
}

// like javascript where we can accept a function 
// as an arugument and the invoke it
func printFunction(fn greeter) {
	s := fn("hello", "world")
	fmt.Println(s)
}

func summaryGetter(converter Converter) {
	p := converter("Plan ID", 5)
	fmt.Println(p)
}
