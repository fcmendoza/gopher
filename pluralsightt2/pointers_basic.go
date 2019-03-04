// As a reference (no pun intented), this is how by ref and by value is done in C#:
// https://dotnetfiddle.net/X16crq

package main

import "fmt"

func main() {
	name := "Jon"
	fmt.Println("\nValue of *name* is", name)

	fmt.Println("Memory value of *name* is", &name)

	ptr := &name
	fmt.Println("Pointer (ptr) to *name* address is", ptr)

	var pun = &name
	fmt.Println("Pointer (pun) to *name* address is", pun)

	var poi *string = &name // the type *string is optional here
	fmt.Println("Pointer (poi) to *name* address is", poi)

	someByRef(&name)
}

func someByRef(some *string) {
	fmt.Println("Value of *some* is", some, "and its value is", *some)
}

/* ----Output----

~/go/src/gopher/pluralsightt2 $ go run pointers_basic.go

Value of *name* is Jon
Memory value of *name* is 0xc000090030
Pointer (ptr) to *name* address is 0xc000090030
Pointer (pun) to *name* address is 0xc000090030
Pointer (poi) to *name* address is 0xc000090030
Value of *some* is 0xc000090030 and its value is Jon

*/
