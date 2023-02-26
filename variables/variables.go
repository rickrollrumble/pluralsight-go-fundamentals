package main

import (
	"fmt"
	"reflect"
	"strconv"
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
var (
	name   = "Rahul R"
	course = "Getting started w Golang"
	// module = 4
	clip   = 2
	module = "4"
)

func main() {
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Name and course are set to", module, "and", clip, ".")
	// types of variables can be checked using runtime reflection and using the reflect package
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	// strings can be converted to integers using type conversion
	// they can then be added. otherwise, data of different types cannot be added.

	// module is passed by value. Behind the scenes, go makes a copy of module sends to Atoi
	// form of immutability, any changes made to copy don't affect original  variable.

	intModule, err := strconv.Atoi(module)
	if err == nil {
		total := intModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	// can get around passing by value using pointers.
	// & returns the memory address of a variable
	fmt.Println("Memory address of *course* variable is", &course)

	// asterisk before type makes ptr a pointer variable
	// assignment says point it to memory location of course
	// so ptr has the hex address
	var ptr *string = &course
	// *ptr is deferencing, so it is the value of the vriable at the location.
	fmt.Println("Pointing to course variable at", ptr, "which holds the value", *ptr)
}
