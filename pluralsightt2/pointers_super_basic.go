// As a reference (no pun intented), this is how by ref and by value is done in C#:
// https://dotnetfiddle.net/X16crq

package main

import "fmt"

func main() {
	name := "Jon"
	fmt.Println("\nValue of *name* is", name)

	fmt.Println("Memory value of *name* is", &name)

	ptr := &name
	fmt.Println("Pointer to *name* address is", ptr)
}

/* ----Output----

~/go/src/gopher/pluralsightt2 $ go run pointers_super_basic.go

Value of *name* is Jon
Memory value of *name* is 0xc000090030
Pointer to *name* address is 0xc000090030

*/
