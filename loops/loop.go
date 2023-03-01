package main

import (
	"fmt"
	"time"
)

func main() {
	rangeLoop()
	nestedLoop()
}

func standardLoop() {
	// timer := 10 is pre-statement
	// timer-- evaluated after each iteration
	// then expression checked
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom")
			break // breaks loop, further iterations do not happen
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}

func infiniteLoop() {
	// for without an expression creates an infinite loop
	for {
	}
}

func rangeLoop() {
	// for range iterates over a list (one entry per iteration of the list)
	coursesInProgress := []string{"English", "Math", "History", "Science"}
	for _, i := range coursesInProgress {
		// underscore blank identifier when not interested in value
		fmt.Println(i)
	}
}

func nestedLoop() {
	coursesInProgress := []string{"English", "Math", "History", "Science"}
	coursesCompleted := []string{"History", "Math"}
	for _, i := range coursesInProgress {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nHey we found a clash.", i, "is in both lists.")
				// continue skips code below it.
				continue
			}
			fmt.Println(i, j)
		}
	}
}

func preStatementLoop() {
	// i:=0; i < 10 are evaluated before the loop
	// i++ evaluated after an iteration
	for i := 0; i < 10; i++ {
	}
}
