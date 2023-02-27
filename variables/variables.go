package main

import (
	"fmt"
	"os"
)

// declaring variables outside main must be done using the var keyword
// go doesn't allow non-declaration statements at the package level
// names have to start with a letter, cannot be a keyboard
// use camelCase
// var (
// 	name, course string
// 	module, clip int
// )

// type inference can be used to write shorter code
// this type of variable declaration makes them available in the package

func variables_main() {
	// name := "Rahul R"
	// updating using environment variable
	name := os.Getenv("USER")
	course := "Getting started w Golang"
	fmt.Println("\nHi", name, "your current course is", course)
	// call function
	updateCourse(course)
	// original value was unchanged, updateCourse made update to a copy of course
	fmt.Println("You're currently watching", course)
	hardUpdateCourse(&course)
	fmt.Println("You're currently watching", course)
}

// Pass by value
func updateCourse(course string) string {
	// change input and return changed input
	// equals operator assigns new value to existing one
	course = "getting started with Docker"
	fmt.Println("Updated course to", course)
	return course
}

func hardUpdateCourse(course *string) string {
	// assign this new value to the location in memory that course pointer variable is reference
	*course = "getting started with Docker"
	fmt.Println("Update course to", *course)
	return *course
}
