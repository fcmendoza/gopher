package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Test from", runtime.GOOS)
}
