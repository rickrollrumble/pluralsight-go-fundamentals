// package name as main is special, run as application
package main

import (
	"fmt"
)

// application won't run without a main function and main package
// not neccessary if writing a shared library, but an executable needs one
func main() {
	fmt.Println("hello world")
}
