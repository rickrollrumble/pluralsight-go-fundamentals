package main

import (
	"fmt"
	"math/rand"
	"time"
)

func switch_case() {
	// any variables declared before the expression are scoped within the switch block
	// unlike other languages, if a switch case is matched, all the others are skipped.
	// breaks are by default, fallthroughs need to be inserted
	// should exactly match condition and type
	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("case 1. kubernetes with lower case 'k'")
	case "Kubernetes":
		fmt.Println("case 2. Kubernetes with upper case 'K")
		fallthrough // only the next case runs.
		// every case has an implicit break, so since the next case din't have a fallthrough, it broke.
	case "k8s":
		fmt.Println("case 3. Kubernetes short name kates")
	case "Istio":
		fmt.Println("case 4. Service mesh is the future")
	default:
		fmt.Println("hit the default")
	}

	// also, it's idiomatic or good practice to match multiple values in a case instead of multiple case matches
	// and not use fallthrough
	switch tempNum := random(); tempNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("we got even number", tempNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("we got odd number", tempNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
