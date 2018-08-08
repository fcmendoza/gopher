package main

import (
	"fmt"
)

func main() {
	// Variables to store course rankings
	//firstRank := "39"   // Docker Deep Dive
	//secondRank := "614" // Docker Clustering

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\n First course is doing better" + 
			" than second course.")
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear ... your first course must" + 
			" be doing abismally!")
	} else {
		fmt.Println("\nBoth courses are either the same" +
			" or something werid is going on.")
	}
}
