// maps are key-value pairs, Go's implementation of hash table
package main

import "fmt"

// to define maps, map[<key type>]<value type>
// key has to be a comparable type (works with == and !=)
// ints, strings, bools, etc. will work.
// key entries HAVE to be unique.

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 0o4
	recentHead2HeadWins := map[string]int{
		"Sunderland": 6,
		"Newcastle":  0,
	}
	fmt.Println(leagueTitles) // Newcastle returned first because Go returns as sorted
	fmt.Println(recentHead2HeadWins)

	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	for i := 0; i < 10; i++ {
		for key, val := range testMap {
			// go returns entire map in sorted order, but iterations can be random.
			// values may be fetched in any order for each iteration.
			fmt.Printf("key is %s, value is %d\n", key, val)
		}
		fmt.Println("End of interation")
	}

	fmt.Println(testMap)
	fmt.Println(testMap["C"])
	// to update values,
	testMap["C"] = 10
	fmt.Println(testMap)
	// if key exists, then the value gets updated. If it doesn't exist, it gets inserted.
	testMap["F"] = 11
	fmt.Println(testMap)
	// to delete a value
	delete(testMap, "C")
	fmt.Println(testMap)
	// maps are reference types, passed to function as a reference. when passed to a function, underlying map is changed.
	// considered good practice to declare size of map while initializing
	// maps are not thread safe. cannot be predicted what happens when written to simultaneously.
}
