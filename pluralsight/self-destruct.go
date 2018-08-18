package main

import (
	"fmt"
	"time"
)

func main() {
    for timer := 10; timer >= 0; timer-- {
    	// if timer == 0 {
    	// 	fmt.Println("Boom!")
    	// 	break
    	// }
    	if timer % 2 == 0 {
    		continue // skips current loop execution. it will skip printing even numbers; 
    			// that is, it will only print odd numbers (9, 7, 5, etc.)
    	}

    	fmt.Println(timer)
    	time.Sleep(1 * time.Second)
    }
}