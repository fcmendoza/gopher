package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta {
		Author: "Nigel Poulton",
		Level: "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("Docker Deep Dive author is:", DockerDeepDive.Author)

	DockerDeepDive.Rating = 1
	fmt.Println("Docker Deep Dive rating is:", DockerDeepDive.Rating)
}
