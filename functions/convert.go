package main

import (
	"fmt"
	"strings"
)

func main() {
	// implementing the titleCase method

	// declaring and initializing two lower case strings
	author := "rahul ramachandran"
	course := "getting started with kubernetes"

	// passing to converter
	// passing by value, so value of author assigned to s1, value of course to s2
	fmt.Println(converter(author, course))
	// Output = RAHUL RAMACHANDRAN Getting Started With Kubernetes

	// declaring a variable that determines Max Verstappen's finishes in formula 1
	// championshipFinishes is a variadic function
	bestFinish := championshipFinishes(12, 5, 6, 4, 3, 3)
	fmt.Println(bestFinish)
}

func converter(s1, s2 string) (str1, str2 string) {
	// input arguments name can be anything
	// receive input strings, return two output strings
	s1 = strings.ToUpper(s1)
	s2 = strings.Title(s2)
	return s1, s2
}

func championshipFinishes(finishes ...int) int {
	// finishes get stored as a slice of ints
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
