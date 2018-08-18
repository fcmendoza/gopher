package main

import (
	"fmt"
)

func main() {
    courseInProg := []string{ "Docker Deep Dive", 
        "Docker Clustering", "Docker and Kubernetes" }

    coursesCompleted := []string { "Docker Deep Dive", 
        "Go Fundamentals", "Puppet Fundamentals" }

    for _, i := range courseInProg {
        for _, j := range coursesCompleted {
            if i == j {
                fmt.Println("\nHey we found a clash.", i, "is in both lists.")
            }
        }
    }
}