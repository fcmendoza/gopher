package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// switch "docker" {
	// 	case "linux": 
	// 		fmt.Println("\nHere are some recommended" + 
	// 			" Linux courses...")
	// 	case "docker": 
	// 		fmt.Println("\nHere are some recommended" + 
	// 			" Docker courses...")
	// 		fallthrough
	// 	case "windows": 
	// 		fmt.Println("\nHere are some recommended" + 
	// 			" Windows courses...")
	// 	default:
	// 		fmt.Println("\nSorry we could not find a match," + 
	// 			" why not try our Top 100 list!")
	// }

	switch tmpNum := random(); tmpNum {
		case 0, 2, 4, 6, 8: 
			fmt.Println("We got an even number -", tmpNum)
		case 1, 3, 5, 7, 9:
			fmt.Println("We got an odd number -", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
