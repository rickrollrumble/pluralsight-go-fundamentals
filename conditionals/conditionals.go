package main

import "fmt"

func main() {
	// dddLengthMins := 275 // docker deep dive course length
	// cawLengthMins := 30 // containers on aws wavelength course length
	// dddLengthMins := 60  // docker deep dive course length
	// cawLengthMins := 120 // containers on aws wavelength course length
	// if dddLengthMins > cawLengthMins {
	// 	fmt.Println("Docker deep dive is longer than Containers on AWS wavelength")
	// } else if dddLengthMins == cawLengthMins {
	// 	fmt.Println("Both courses are the same length")
	// } else {
	// 	fmt.Println("Containers on AWS wavelength is longer")
	// }

	// if blocks also allow simple initialization statements before expression evaluation
	// a reason to do this is scoped without boundaries of if block, then get garbage collected
	// if any conditional is true, the other conditions are skipped
	// unlimited else ifs are possible, but only one final else.
	if dddLengthMins, cawLengthMins := 275, 30; dddLengthMins > cawLengthMins {
		fmt.Println("Docker deep dive is longer than Containers on AWS wavelength")
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Both courses are the same length")
	} else {
		fmt.Println("Containers on AWS wavelength is longer")
	}

	if dddLengthMins, cawLengthMins := 275, 30; dddLengthMins > cawLengthMins {
		fmt.Println("Docker deep dive is longer than Containers on AWS wavelength")
		if dddLengthMins > 240 {
			fmt.Println("but it's so long it'll put viewers to sleep")
		}
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Both courses are the same length")
	} else {
		fmt.Println("Containers on AWS wavelength is longer")
	}

	// refer switch.go
	switch_case()
}
