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

func main() {
	name = "Rahul R"
	course = "Getting started w Golang"
	// module = 4
	clip = 2
	module = "4"

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Name and course are set to", module, "and", clip, ".")
	// types of variables can be checked using runtime reflection and using the reflect package
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	// strings can be converted to integers using type conversion
	// they can then be added. otherwise, data of different types cannot be added.
	intModule, err := strconv.Atoi(module)
	if err == nil {
		total := intModule + clip
		fmt.Println("Module plus clip equals", total)
	}
}
