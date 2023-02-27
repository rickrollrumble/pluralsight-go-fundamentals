package main

import "fmt"

func consts_main() {
	// no shorthand way of declaring constants like variables
	// this variable cannot be changed
	const c = 186000 // speed of light
	fmt.Println("the speed of light is", c, "miles per second")
}
